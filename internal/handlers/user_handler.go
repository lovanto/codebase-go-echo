package handlers

import (
	"codebase-go-echo/internal/services"
	"codebase-go-echo/pkg/utils"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	// Set a context with a timeout
	ctx := c.Request().Context()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Parse query parameters for pagination
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10 // Default limit
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0 // Default offset
	}

	// Fetch users with pagination
	users, err := services.GetUsers(timeoutCtx, limit, offset)
	if err != nil {
		response := utils.ErrorResponse(http.StatusInternalServerError, "Failed to fetch users", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := utils.SuccessResponse("Users fetched successfully", users)
	return c.JSON(http.StatusOK, response)
}

func GetUsersPaginated(c echo.Context) error {
	ctx := c.Request().Context()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	users, totalCount, err := services.GetUsersPaginated(timeoutCtx, limit, offset)
	if err != nil {
		response := utils.ErrorResponse(http.StatusInternalServerError, "Failed to fetch users", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	totalPages := (totalCount + limit - 1) / limit

	response := utils.PaginatedResponse("Users fetched successfully", users, page, limit, totalPages, totalCount)
	return c.JSON(http.StatusOK, response)
}
