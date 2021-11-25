package middleware

import (
	"bytes"
	"data-pipeline/config"
	"data-pipeline/service"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// This middleware function will run before the response is returned
func ErrorNotify() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//prepare buffer for writing response body
			resBody := new(bytes.Buffer)

			mw := io.MultiWriter(c.Response().Writer, resBody)
			writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
			c.Response().Writer = writer
			defer func() {
				var responseBody string

				if c.Response().Size < 50000 {
					responseBody = resBody.String()
				}

				statusCode := c.Response().Status

				if statusCode > 399 {
					service.SendMail(config.Email, config.Email, "Data Pipeline Job Failed", responseBody)
				}
			}()

			return next(c)
		}
	}
}

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
