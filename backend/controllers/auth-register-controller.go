package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	. "github.com/c0de-runn3r/libraryManagementSystem/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/c0de-runn3r/libraryManagementSystem/db/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func HandleRegister(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	user := new(models.User)
	json.NewDecoder(c.Request().Body).Decode(&user)

	if user.Name == "" || user.Surname == "" { //TODO - make validator via echo.validate
		return c.JSON(http.StatusBadRequest, "not enough data")
	}

	database.Create(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusConflict, "user already exists")
	}

	Log("debug", "Registered new user")
	return c.JSON(http.StatusOK, "user succesfully registered")
}

func HandleLogin(c echo.Context) error {
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

	expTime := time.Now().Add(time.Hour * 24)
	token, err := GenerateToken(user, expTime, []byte(GetJWTSecret()))
	if err != nil {
		Log("error", fmt.Sprintln("could not generate jwt token: %w", err))
		return c.JSON(http.StatusInternalServerError, "could not login")
	}
	SetTokenCookie(token, expTime, c)

	Log("debug", "New login")
	return c.JSON(http.StatusOK, "success")
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET_KEY")
}

func GenerateToken(user *models.User, expTime time.Time, secret []byte) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: expTime.Unix(),
	})

	token, err := claims.SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func SetTokenCookie(token string, expTime time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = expTime
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func Skipper(c echo.Context) bool {
	database := c.Get(dbContextKey).(*gorm.DB)
	var user models.User
	cookie, _ := c.Cookie("jwt")
	if cookie == nil {
		return false
	}
	id, err := getIDbyJWT(cookie.Value)
	if err != nil {
		return false
	}
	database.Table("users").Where("id = ?", id).First(&user)
	if user.ID == 0 {
		return false
	}

	return true
}

func JWTErrorChecker(err error) error {
	return echo.NewHTTPError(http.StatusUnauthorized, "Unuthorized")
}
