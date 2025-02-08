package handlers

import (
	"codebase-go-echo/internal/models"
	"codebase-go-echo/internal/services"
	"codebase-go-echo/pkg/utils"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	ctx := c.Request().Context()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	users, err := services.GetUsers(timeoutCtx, limit, offset)
	if err != nil {
		response := utils.ErrorResponse(http.StatusInternalServerError, "failed to fetch users", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := utils.SuccessResponse("users fetched successfully", users)
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
		response := utils.ErrorResponse(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	totalPages := (totalCount + limit - 1) / limit

	response := utils.PaginatedResponse("users fetched successfully", users, page, limit, totalPages, totalCount)
	return c.JSON(http.StatusOK, response)
}

func CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	var user models.User
	if err := c.Bind(&user); err != nil {
		response := utils.ErrorResponse(http.StatusBadRequest, "Invalid input", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := services.CreateUser(ctx, &user); err != nil {
		response := utils.ErrorResponse(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	response := utils.SuccessResponse("User created successfully", userResponse)
	return c.JSON(http.StatusCreated, response)
}

func UpdateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ErrorResponse(http.StatusBadRequest, "invalid user ID", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		response := utils.ErrorResponse(http.StatusBadRequest, "invalid input", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := services.UpdateUser(ctx, id, &user); err != nil {
		response := utils.ErrorResponse(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	response := utils.SuccessResponse("User updated successfully", userResponse)
	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ErrorResponse(http.StatusBadRequest, "invalid user ID", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := services.DeleteUser(ctx, id); err != nil {
		response := utils.ErrorResponse(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := utils.SuccessResponse("User deleted successfully", nil)
	return c.JSON(http.StatusOK, response)
}
