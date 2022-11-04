package router

import (
	middlewares "golang-echo-framework/src/core"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	//set all middlewares
	middlewares.SetMainMiddleWares(e)

	return e
}
