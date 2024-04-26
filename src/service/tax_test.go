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

	t.Run("Cal Tax less than 150000 ", func(t *testing.T) {
		want := TaxCalculationResponse{
			Tax: 0,
		}
		got, err := TaxCalculation(TaxCalculationRequest{
			TotalIncome: 150000.0,
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

	t.Run("Cal Tax less than 1000000 ", func(t *testing.T) {
		want := TaxCalculationResponse{
			Tax: 66000,
		}
		got, err := TaxCalculation(TaxCalculationRequest{
			TotalIncome: 1000000.0,
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

	t.Run("Cal Tax less than 2000000 ", func(t *testing.T) {
		want := TaxCalculationResponse{
			Tax: 188000,
		}
		got, err := TaxCalculation(TaxCalculationRequest{
			TotalIncome: 2000000.0,
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

	// t.Run("Cal tax more than 5000000", func (t *testing.T) {
	// 	want := TaxCalculationResponse{
	// 		Tax: 20790000,
	// 	}
	// 	got, err := TaxCalculation(TaxCalculationRequest{
	// 		TotalIncome: 8000000.0,
	// 		Wht: 0,
	// 		Allowances: []Allowance{
	// 			{
	// 				AllowanceType: "donation",
	// 				Amount: 0,
	// 			},
	// 		},
	// 	})

	// 	assert.Nil(t, err)
	// 	assert.NotNil(t, got)

	// 	assert.Equal(t, want, *got)
	// })
}

func TestCalculateTaxWithWHT(t *testing.T) {
	t.Run("Cal Tax with WHT", func(t *testing.T) {
		want := TaxCalculationResponse{
			Tax: 4000,
		}
		got, err := TaxCalculation(TaxCalculationRequest{
			TotalIncome: 500000.0,
			Wht: 25000.0,
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

	t.Run("Cal Tax with WHT less than 150000 ", func(t *testing.T) {
		want := TaxCalculationResponse{
			Tax: 0,
		}
		got, err := TaxCalculation(TaxCalculationRequest{
			TotalIncome: 150000.0,
			Wht: 25000.0,
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

	t.Run("Cal Tax with WHT less than 1000000 ", func(t *testing.T) {
		want := TaxCalculationResponse{
			Tax: 41000,
		}
		got, err := TaxCalculation(TaxCalculationRequest{
			TotalIncome: 1000000.0,
			Wht: 25000.0,
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

func TestCalculateTaxWithAllowances(t *testing.T) {
	t.Run("Cal Tax with donation", func(t *testing.T) {
		want := TaxCalculationResponse{
			Tax: 19000.0,
		}
		got, err := TaxCalculation(TaxCalculationRequest{
			TotalIncome: 500000.0,
			Wht: 0,
			Allowances: []Allowance{
				{
					AllowanceType: "donation",
					Amount: 200000.0,
				},
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, got)

		assert.Equal(t, want, *got)
	})

	t.Run("Cal Tax with k-receipt", func(t *testing.T) {
		want := TaxCalculationResponse{
			Tax: 14000.0,
		}

		got, err := TaxCalculation(TaxCalculationRequest{
			TotalIncome: 500000.0,
			Wht: 0,
			Allowances: []Allowance{
				{
					AllowanceType: "k-receipt",
					Amount: 200000.0,
				},
				{
					AllowanceType: "donation",
					Amount: 100000.0,
				},
			
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, got)

		assert.Equal(t, want, *got)
	})
}
