package service

import "fmt"
import "math"


// import "errors"

func TaxCalculation(request TaxCalculationRequest) (*TaxCalculationResponse, error) {

	netIncome := request.TotalIncome

	tax := TaxCalculationResponse{}

	// allowances
	netIncome -= 60000
	for _, allowance := range request.Allowances {
		if allowance.AllowanceType == "donation" {
			netIncome -= min(allowance.Amount, 100000.0)
		} else if allowance.AllowanceType == "k-receipt" {
			netIncome -= min(allowance.Amount, 50000.0)
		}
	}

	tax.TaxLevel = []TaxLevel{
		{
			Level: "0-150,000",
			Tax:   0,
		},
		{
			Level: "150,001-500,000",
			Tax:   0,
		},
		{
			Level: "500,001-1,000,000",
			Tax:   0,
		},
		{
			Level: "1,000,001-2,000,000",
			Tax:   0,
		},
		{
			Level: "2,000,001 ขึ้นไป",
			Tax:   0,
		},
	}

	taxCalculaled := 0.0
	if netIncome > 0 {
		tax.TaxLevel[0].Tax += 0;
	}
	if netIncome > 150000 {
		tax.TaxLevel[1].Tax += 0.1 * (min(netIncome, 500000) - 150000)
	}
	if netIncome > 500000 {
		tax.TaxLevel[2].Tax += 0.15 * (min(netIncome, 1000000) - 500000)
	}
	if netIncome > 1000000 {
		tax.TaxLevel[3].Tax += 0.2 * (min(netIncome, 2000000) - 1000000)
	}
	if netIncome > 2000000 {
		tax.TaxLevel[4].Tax += 0.35 * (netIncome - 2000000)
	}

	for _, taxLevel := range tax.TaxLevel {
		taxCalculaled += taxLevel.Tax
	}

	// wht
	taxCalculaled -= request.Wht

	if taxCalculaled < 0 {
		taxCalculaled = 0
	}

	tax.Tax = math.Round(taxCalculaled*100)/100

	fmt.Println(tax)

	return &tax, nil
}
