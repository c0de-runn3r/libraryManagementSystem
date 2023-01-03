package controllers

import (
	. "lms/utils"
	"net/http"
	"os"
	"strconv"
	"time"

	"lms/db/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const dbContextKey = "__db" // just for dbMiddleware use. See below

// Middleware for echo package to pass the database into API endpoints' handlers
func dbMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(dbContextKey, db)
			return next(c)
		}
	}
}

func AttachUsersController(g *echo.Group, db *gorm.DB) {

	Log("info", "Attaching USERS controller.")

	g.Use(dbMiddleware(db))

	g.GET("/get-user", handleGetUser)
	g.POST("/register", handleRegister)
	g.POST("/login", handleLogin)
	g.POST("/logout", handleLogout)
}

func handleGetUser(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)
	cookie, _ := c.Cookie("jwt")
	if cookie == nil {
		Log("error", "Can't get user: unauthenticated (no cookie set)")
		return c.JSON(http.StatusUnauthorized, "unauthenticated")
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		Log("error", "Can't get user: unauthenticated (cookie doesn't match)")
		return c.JSON(http.StatusUnauthorized, "unauthenticated")
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.Where("id = ?", claims.Issuer).First(&user)

	Log("debug", "Handled get user")
	return c.JSON(http.StatusOK, user)
}

func handleRegister(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	password, _ := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), 14)

	user := models.User{
		Name:     c.FormValue("name"),
		Surname:  c.FormValue("surname"),
		Email:    c.FormValue("email"),
		Password: password,
	}

	database.Create(&user)

	Log("debug", "Registered new user")
	return c.JSON(http.StatusOK, user)
}

func handleLogin(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)
	user := models.User{
		Email: c.FormValue("email"),
	}

	database.Where("email = ?", user.Email).First(&user)
	if user.ID == 0 { //If the ID return is '0' then there is no such email present in the DB
		Log("info", "could not login: user not found")
		return c.JSON(http.StatusNotFound, "user not found")
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(c.FormValue("password"))); err != nil {
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
