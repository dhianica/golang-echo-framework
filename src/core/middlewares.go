package middlewares

import (
	"github.com/labstack/echo"
)

func SetMainMiddleWares(e *echo.Echo) {
	e.Use(serverHeader)
}

//Custom Middleware
// ServerHeader middleware adds a `Server` header to the response.
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Custom-Header", "blah!!!")
		return next(c)
	}
}
