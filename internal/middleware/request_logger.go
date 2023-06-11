package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"regexp"
	"social_media/pkg/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

type responseBodyWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (mw *MiddlewareManager) RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		re := regexp.MustCompile(`\r|\n|\x{0020}`)

		var bodyAsByteArray []byte
		if c.Request().Body != nil {
			bodyAsByteArray, _ = ioutil.ReadAll(c.Request().Body)
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyAsByteArray))
		}

		rbw := &responseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Response().Writer}
		c.Response().Writer = rbw

		err := next(c)

		req := c.Request()
		requestID := utils.GetRequestID(c)

		mw.logger.Infof("REQUEST-ID: %s, METHOD: %s, URI: %s, STATUS: %d, REQUEST: %s, RESPONSE: %s",
			requestID, req.Method, req.URL, c.Response().Status, re.ReplaceAllString(strings.TrimSpace(string(bodyAsByteArray)), ""), re.ReplaceAllString(rbw.body.String(), ""),
		)

		return err
	}
}
