package main

import (
	"codebase-go-echo/config"
	"codebase-go-echo/internal/routes"
	"codebase-go-echo/pkg/databases/postgresql"
	internal_middleware "codebase-go-echo/pkg/middlewares"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()
	if cfg.PostgreSqlDsn == "" || cfg.ServerPort == "" {
		slog.Error("Missing required environment variables")
		os.Exit(1)
	}

	postgresql.ConnectDB(cfg.PostgreSqlDsn)
	postgresql.SetMaxConnections(50, 25, 30*time.Minute)
	defer postgresql.CloseDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(internal_middleware.SecurityConfig())
	e.Use(internal_middleware.RateLimitConfig())
	e.Use(internal_middleware.CORSConfig())

	routes.RegisterRoutes(e)

	port := fmt.Sprintf(":%s", cfg.ServerPort)
	slog.Info("Server running", "port", port)

	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			slog.Error("Server error", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		slog.Error("Error shutting down", "error", err)
	}
}
