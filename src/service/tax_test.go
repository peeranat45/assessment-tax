package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxWithoutWHT(t *testing.T) {
	t.Run("Cal Tax without donation", func(t *testing.T) {
		want := TaxCalculationResponse{
			Tax: 29000,
		}
		got, err := TaxCalculation(TaxCalculationRequest{
			TotalIncome: 500000.0,
			Wht: 0,
			Allowances: []Allowance{
				{
					AllowanceType: "donation",
					Amount: 0,
				},
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, got)

		assert.Equal(t, want, *got)

	})
}