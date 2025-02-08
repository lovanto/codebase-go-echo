package main

import (
	"codebase-go-echo/config"
	"codebase-go-echo/internal/routes"
	"codebase-go-echo/pkg/databases/postgresql"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()
	postgresql.ConnectDB(cfg.PostgreSqlDsn)
	postgresql.SetMaxConnections(50, 25, time.Minute*30)

	defer postgresql.CloseDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e)

	port := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server running on %s", port)
	e.Start(port)
}
