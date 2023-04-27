package presenter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SuccessResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, errorResponse{
		Code:    http.StatusOK,
		Message: message,
	})
}

func ErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, errorResponse{
		Code:    code,
		Message: message,
	})
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
