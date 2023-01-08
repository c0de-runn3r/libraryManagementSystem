package controllers

import (
	"github.com/labstack/echo"
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
