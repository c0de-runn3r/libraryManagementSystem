package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	. "github.com/c0de-runn3r/libraryManagementSystem/utils"

	"github.com/c0de-runn3r/libraryManagementSystem/db/models"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func AttachUsersController(g *echo.Group, db *gorm.DB) {

	Log("info", "Attaching USERS controller.")

	g.POST("/new-user", handleAddNewUser)

	g.GET("/get-all-users", handleGetAllUsers)
}

func handleAddNewUser(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	user := new(models.User)
	json.NewDecoder(c.Request().Body).Decode(&user)

	if user.Name == "" { //TODO - make validator via echo.validate
		return c.JSON(http.StatusBadRequest, "not enough data")
	}

	user.Name = ConvertToTitleCase(user.Name)

	database.Create(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusConflict, "user already exists")
	}

	Log("debug", "Registered new user")
	return c.JSON(http.StatusOK, "user succesfully registered")
}

func handleGetAllUsers(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	var users []*models.User
	database.Table("users").Find(&users)

	Log("debug", "Handled get all users")
	return c.JSON(http.StatusOK, users)
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

// NOT USED

// func handleRegisterByUser(c echo.Context) error {
// 	database := c.Get(dbContextKey).(*gorm.DB)

// 	user := new(models.User)
// 	json.NewDecoder(c.Request().Body).Decode(&user)

// 	if len(string(user.Password)) < 8 {
// 		return c.JSON(http.StatusBadRequest, "password is too short")
// 	}
// 	if user.Name == "" || user.Surname == "" || user.Email == "" { //TODO - make validator via echo.validate
// 		return c.JSON(http.StatusBadRequest, "not enough data")
// 	}

// 	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

// 	user.Password = string(password)

// 	user.VerificationCode = generateVerificationCode(30)

// 	database.Create(&user)
// 	if user.ID == 0 {
// 		return c.JSON(http.StatusConflict, "user already exists")
// 	}

// 	err := SendVerificationCodeViaEmail(user.Email, user.VerificationCode)
// 	if err != nil {
// 		Log("error", fmt.Sprintf("error sending confirmation code via email: %e", err))
// 		database.Delete(&user)
// 		return c.JSON(http.StatusOK, "error sending confirmation code")
// 	}

// 	message := "We sent an email with a verification code to " + user.Email
// 	Log("debug", "Registered new user")
// 	return c.JSON(http.StatusOK, message)
// }

// func handleVerifyEmail(c echo.Context) error {
// 	database := c.Get(dbContextKey).(*gorm.DB)

// 	code := c.Param("code")

// 	var user models.User

// 	res := database.Where("verification_code = ?", code).First(&user)
// 	if res.Error != nil {
// 		return c.JSON(http.StatusForbidden, "invalid verification code or user doesn't exists")
// 	}
// 	if user.EmailVerified {
// 		return c.JSON(http.StatusForbidden, "email already verified")
// 	}

// 	user.EmailVerified = true
// 	database.Save(&user)

// 	return c.JSON(http.StatusOK, "email verified successfully")
// }
