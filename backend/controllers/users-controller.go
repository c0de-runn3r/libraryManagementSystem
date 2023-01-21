package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	. "github.com/c0de-runn3r/libraryManagementSystem/utils"

	"github.com/c0de-runn3r/libraryManagementSystem/db/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AttachUsersController(g *echo.Group, db *gorm.DB) {

	Log("info", "Attaching USERS controller.")

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:      Skipper,
		Claims:       &jwt.StandardClaims{},
		SigningKey:   []byte(GetJWTSecret()),
		TokenLookup:  "cookie:jwt", // "<source>:<name>"
		ErrorHandler: middleware.JWTErrorHandler(JWTErrorChecker),
	}))

	g.GET("/get-user", handleGetUser)
	g.GET("/get-all-users", handleGetAllUsers)
	g.GET("/verifyemail/:code", handleVerifyEmail)
	g.POST("/logout", handleLogout)
}

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
func handleGetAllUsers(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	user, err := GetUserByRequestContext(c, database)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}
	if user.Role != models.Librarian && user.Role != models.Manager {
		return c.JSON(http.StatusForbidden, "you have no roots to preform this request")
	}

	var users []*models.UserResponse
	database.Table("users").Find(&users)

	Log("debug", "Handled get all users")
	return c.JSON(http.StatusOK, users)
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
		return []byte(GetJWTSecret()), nil
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

// Not used

func handleRegisterByUser(c echo.Context) error {
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

	err := SendVerificationCodeViaEmail(user.Email, user.VerificationCode)
	if err != nil {
		Log("error", fmt.Sprintf("error sending confirmation code via email: %e", err))
		database.Delete(&user)
		return c.JSON(http.StatusOK, "error sending confirmation code")
	}

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
