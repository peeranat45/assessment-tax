package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"

	"github.com/KKGo-Software-engineering/assessment-tax/src/service"
)

func TestCalculateTaxHandler(t *testing.T) {
	t.Run("Cal Tax without WHT", func(t *testing.T) {
		e := echo.New()
		e.POST("tax/calculations", CalculateTax)

		body := service.TaxCalculationRequest{
			TotalIncome: 500000.0,
			Wht:         0,
			Allowances: []service.Allowance{
				{
					AllowanceType: "donation",
					Amount:        0,
				},
			},
		}

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(body)

		req := httptest.NewRequest(http.MethodPost, "/tax/calculations", &buf)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		CalculateTax(c)

		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("Cal Tax with WHT", func(t *testing.T) {
		e := echo.New()
		e.POST("tax/calculations", CalculateTax)

		body := service.TaxCalculationRequest{
			TotalIncome: 500000.0,
			Wht:         25000.0,
			Allowances: []service.Allowance{
				{
					AllowanceType: "donation",
					Amount:        0,
				},
			},
		}

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(body)

		req := httptest.NewRequest(http.MethodPost, "/tax/calculations", &buf)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		CalculateTax(c)

	})
}
