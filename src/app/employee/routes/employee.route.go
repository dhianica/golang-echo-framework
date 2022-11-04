package api

import (
	employee "golang-echo-framework/src/app/employee/handlers"

	"github.com/labstack/echo"
)

func routeEmployee(e *echo.Echo) {
	e.GET("/employee/:data", employee.Getemployee)
}
