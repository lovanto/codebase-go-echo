package main

import (
	"codebase-go-echo/config"
	"codebase-go-echo/internal/handlers"
	"codebase-go-echo/pkg/databases/postgresql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()
	postgresql.ConnectDB(cfg.PostgreSqlDsn)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.GetHealthCheck)
	e.GET("/users", handlers.GetUsers)

	port := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server running on %s", port)
	e.Start(port)
}
