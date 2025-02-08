package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SecurityConfig returns security middleware configuration.
func SecurityConfig() echo.MiddlewareFunc {
	return middleware.Secure()
}
