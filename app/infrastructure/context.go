package infrastructure

import (
	"errors"
	"go-echo-sample/app/interface/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}

func (cc *CustomContext) Bind(i interface{}) error {
	err := cc.Context.Bind(i)
	if err != nil {
		return &HTTPError{err}
	}

	return nil
}

type HTTPError struct {
	error
}

func (cc *CustomContext) CustomResponse() controllers.Response {
	return cc.Context.Response()
}

func (e *HTTPError) Status() int {
	ie := e.error
	if ie == nil {
		return http.StatusOK
	}

	for {
		switch err := ie.(type) {
		case *echo.HTTPError:
			return err.Code
		}

		ie = errors.Unwrap(ie)
		if ie == nil {
			return http.StatusInternalServerError
		}
	}
}

func (e *HTTPError) Unwrap() error {
	return e.error
}

func CustomContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{c}

		return next(cc)
	}
}
