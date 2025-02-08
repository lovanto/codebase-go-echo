package handlers

import (
	"codebase-go-echo/pkg/databases/postgresql"
	"codebase-go-echo/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetHealthCheck(c echo.Context) error {
	err := postgresql.PingDB()
	var response utils.Response
	if err != nil {
		response = utils.ErrorResponse(http.StatusInternalServerError, "Database is not working, please check your database", nil)
	} else {
		response = utils.SuccessResponse("Service is running successfully", nil)
	}
	return c.JSON(http.StatusOK, response)
}
