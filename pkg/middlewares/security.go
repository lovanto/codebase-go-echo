package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SecurityConfig() echo.MiddlewareFunc {
	return middleware.Secure()
}
