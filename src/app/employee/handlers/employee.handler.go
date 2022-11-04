package mngthandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mngtservice "golang-echo-framework/src/app/employee/services"

	"github.com/labstack/echo"
)

//http://localhost:8080/employee/json?name=arnold&type=fluffy
func Getemployee(c echo.Context) error {
	employeeName := c.QueryParam("name")
	employeeType := c.QueryParam("type")
	dataType := c.Param("data")
	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your employee name is : %s\nand employee type is : %s\n", employeeName, employeeType))
	} else if dataType == "json" {
		employee := mngtservice.Employee{
			Name: employeeName,
			Type: employeeType,
		}
		return c.JSON(http.StatusOK, employee)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data type as Sting or JSON",
		})
	}

}

func Addemployee(c echo.Context) error {
	employee := mngtservice.Employee{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&employee)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is yout employee %#v", employee)
	return c.String(http.StatusOK, "We got your employee!!!")
}
