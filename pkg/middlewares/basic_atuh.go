package middleware

import (
	"codebase-go-echo/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicAuthMiddleware() echo.MiddlewareFunc {
	cfg := config.LoadConfig()
	println(cfg.BasicAuthUsername, cfg.BasicAuthPassword)
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == cfg.BasicAuthUsername && password == cfg.BasicAuthPassword {
			return true, nil
		}
		return false, nil
	})
}
