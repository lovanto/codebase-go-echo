package handlers

import (
	"codebase-go-echo/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetHealthCheck(c echo.Context) error {
	response := utils.SuccessResponse("Service is running", nil)
	return c.JSON(http.StatusOK, response)
}
