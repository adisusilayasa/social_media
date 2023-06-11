package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	ResponseCode int         `json:"code"`
	Message      string      `json:"message"`
	Status       string      `json:"status"`
	Data         interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, code int, message string, data interface{}) (int, interface{}) {
	response := Response{
		ResponseCode: code,
		Message:      message,
		Status:       "success",
		Data:         data,
	}

	return code, response
}

func ErrorResponse(c echo.Context, err error) (int, interface{}) {
	statusCode := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		statusCode = he.Code
	}

	response := Response{
		ResponseCode: statusCode,
		Message:      err.Error(),
		Status:       "failed",
		Data:         nil,
	}

	return statusCode, response
}
