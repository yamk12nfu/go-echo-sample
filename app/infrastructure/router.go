package infrastructure

import (
	"go-echo-sample/app/interface/controllers"
	"go-echo-sample/app/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *Router {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	return &Router{e}
}

type Router struct {
	Echo *echo.Echo
}

func (r *Router) Start() {
	r.Echo.Logger.Fatal(r.Echo.Start(":8001"))
}

func HTTPErrorhandler(err error, c echo.Context) {
	status := usecases.GetErrorStatus(err)
	body := map[string]interface{}{"status": status, "error": err.Error()}

	if err := c.JSON(status, body); err != nil {
		c.Logger().Error(err)
	}
}

type ControllerFunc func(controllers.Context) error

func newHandlerFunc(f ControllerFunc) echo.HandlerFunc {
	return func(c echo.Context) error { return f(c.(*CustomContext)) }
}
