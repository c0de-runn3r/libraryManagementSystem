package controllers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

const dbContextKey = "__db" // just for dbMiddleware use. See below

// Middleware for echo package to pass the database into API endpoints' handlers
func DBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(dbContextKey, db)
			return next(c)
		}
	}
}

// Config for JWT middleware

var JWTMiddlewareCustomConfig = middleware.JWTConfig{
	Skipper:      Skipper,
	Claims:       &jwt.StandardClaims{},
	SigningKey:   []byte(GetJWTSecret()),
	TokenLookup:  "cookie:jwt", // "<source>:<name>"
	ErrorHandler: middleware.JWTErrorHandler(JWTErrorChecker),
}
