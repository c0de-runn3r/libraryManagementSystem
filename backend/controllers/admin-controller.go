package controllers

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"gorm.io/gorm"

	"github.com/c0de-runn3r/libraryManagementSystem/db/models"
	. "github.com/c0de-runn3r/libraryManagementSystem/utils"
)

func AttachAdminsController(g *echo.Group, db *gorm.DB) {

	Log("info", "Attaching ADMINS controller.")

	g.GET("/get-admin", handleGetAdmin)
}

func handleGetAdmin(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	user, err := GetAdminByRequestContext(c, database)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	Log("debug", "Handled get admin")
	return c.JSON(http.StatusOK, user)
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

func GetAdminByRequestContext(c echo.Context, database *gorm.DB) (models.Admin, error) {
	var admin models.Admin
	cookie, _ := c.Cookie("jwt")
	if cookie == nil {
		Log("debug", "Can't get admin: unauthenticated (no cookie set)")
		return admin, fmt.Errorf("unauthenticated")
	}
	id, err := getIDbyJWT(cookie.Value)
	if err != nil {
		Log("debug", "Can't get admin: unauthenticated (cookie doesn't match)")
		return admin, fmt.Errorf("unauthenticated")
	}
	database.Table("admins").Where("id = ?", id).First(&admin)
	return admin, nil
}
