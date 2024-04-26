package handler

import (
	"fmt"

	"github.com/KKGo-Software-engineering/assessment-tax/src/service"
	"github.com/labstack/echo/v4"

	"net/http"
)

func CalculateTax(c echo.Context) error {
	request := service.TaxCalculationRequest{}
	err := c.Bind(&request)
	fmt.Println(err)
	if err != nil {
		return c.JSON(400, map[string]string{
			"message": "Invalid request",
		})
	}

	tax, err := service.TaxCalculation(request)
	if err != nil {
		return c.JSON(500, map[string]string{
			"message": "Internal server error",
		})
	}
	return c.JSON(http.StatusCreated, tax)
}