package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

// RequestLogging returns an 'echo' middleware
func RequestLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println(c.Path())
		return next(c)
	}
}
