package utils

import (
	"context"

	"github.com/labstack/echo/v4"
)

func GetRequestID(c echo.Context) string {
	return c.Request().Header.Get(echo.HeaderXRequestID)
}

type ReqIDCtxKey struct{}

func GetRequestCtx(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}

// func GetIPAddress(c *echo.Context) string {
// 	return c.Request().RemoteAddr
// }

func ReadRequest(c echo.Context, request interface{}) error {
	if err := c.Bind(request); err != nil {
		return err
	}
	return c.Validate(request)
}
