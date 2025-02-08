package routes

import (
	"codebase-go-echo/internal/handlers"
	basic_auth "codebase-go-echo/pkg/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")

	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	api.Use(basic_auth.BasicAuthMiddleware())

	rootGroup := api.Group("")
	rootGroup.GET("/health_check", handlers.GetHealthCheck)

	userGroup := api.Group("/users")
	userGroup.GET("/", handlers.GetUsers)
}
