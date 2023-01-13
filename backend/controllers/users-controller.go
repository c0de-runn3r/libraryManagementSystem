package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	. "github.com/c0de-runn3r/libraryManagementSystem/utils"

	"github.com/c0de-runn3r/libraryManagementSystem/db/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AttachUsersController(g *echo.Group, db *gorm.DB) {

	Log("info", "Attaching USERS controller.")

	g.Use(dbMiddleware(db))

	g.GET("/get-user", handleGetUser)
	g.POST("/register", handleRegister)
	g.GET("/verifyemail/:code", handleVerifyEmail)
	g.POST("/login", handleLogin)
	g.POST("/logout", handleLogout)
} // TODO move all database interactios to db package

func handleGetUser(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	user, err := GetUserByRequestContext(c, database)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	userResp := models.UserResponse{
		Name:          user.Name,
		Surname:       user.Surname,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		Role:          user.Role,
	}

	Log("debug", "Handled get user")
	return c.JSON(http.StatusOK, userResp)
}

func handleRegister(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	user := new(models.User)
	json.NewDecoder(c.Request().Body).Decode(&user)

	if len(string(user.Password)) < 8 {
		return c.JSON(http.StatusBadRequest, "password is too short")
	}
	if user.Name == "" || user.Surname == "" || user.Email == "" { //TODO - make validator via echo.validate
		return c.JSON(http.StatusBadRequest, "not enough data")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	user.Password = string(password)

	user.VerificationCode = generateVerificationCode(30)

	database.Create(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusConflict, "user already exists")
	}

	// err := SendVerificationCodeViaEmail(user.Email, user.VerificationCode)
	// if err != nil {
	// 	Log("error", fmt.Sprintf("error sending confirmation code via email: %e", err))
	// 	database.Delete(&user)
	// 	return c.JSON(http.StatusOK, "error sending confirmation code")
	// }

	message := "We sent an email with a verification code to " + user.Email
	Log("debug", "Registered new user")
	return c.JSON(http.StatusOK, message)
}

func handleVerifyEmail(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	code := c.Param("code")

	var user models.User

	res := database.Where("verification_code = ?", code).First(&user)
	if res.Error != nil {
		return c.JSON(http.StatusForbidden, "invalid verification code or user doesn't exists")
	}
	if user.EmailVerified {
		return c.JSON(http.StatusForbidden, "email already verified")
	}

	user.EmailVerified = true
	database.Save(&user)

	return c.JSON(http.StatusOK, "email verified successfully")
}

func handleLogin(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	data := new(models.User)
	json.NewDecoder(c.Request().Body).Decode(&data)

	user := new(models.User)

	database.Where("email = ?", data.Email).First(&user)
	if user.ID == 0 { //If the ID return is '0' then there is no such email present in the DB
		Log("info", "could not login: user not found")
		return c.JSON(http.StatusNotFound, "user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		Log("info", "could not login: incorrect password")
		return c.JSON(http.StatusBadRequest, "incorrect password")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		Log("error", "could not login: jwt secret key error")
		return c.JSON(http.StatusInternalServerError, "could not login")
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	Log("debug", "New login")
	return c.JSON(http.StatusOK, "success")
}

func handleLogout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	Log("debug", "User logout")
	return c.JSON(http.StatusOK, "success")
}

func getIDbyJWT(JWtoken string) (string, error) {
	token, err := jwt.ParseWithClaims(JWtoken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}

func GetUserByRequestContext(c echo.Context, database *gorm.DB) (models.User, error) {
	var user models.User
	cookie, _ := c.Cookie("jwt")
	if cookie == nil {
		Log("debug", "Can't get user: unauthenticated (no cookie set)")
		return user, fmt.Errorf("unauthenticated")
	}
	id, err := getIDbyJWT(cookie.Value)
	if err != nil {
		Log("debug", "Can't get user: unauthenticated (cookie doesn't match)")
		return user, fmt.Errorf("unauthenticated")
	}
	database.Table("users").Where("id = ?", id).First(&user)
	return user, nil
}

func generateVerificationCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "abcdefghijklmnopqrstuvwxyz"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
