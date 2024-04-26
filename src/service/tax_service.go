package service

// import "errors"


func TaxCalculation(request TaxCalculationRequest) (*TaxCalculationResponse, error) {

	taxCalculated := request.TotalIncome
	
	

	

	

	// allowances
	taxCalculated -= 60000
	// for _, allowance := range request.Allowances {
	// 	if allowance.AllowanceType == "donation" {
	// 		taxCalculated -= min(allowance.Amount, 100000.0)
	// 	}
	// }

	if taxCalculated <= 150000 {
		taxCalculated = 0
	}else if taxCalculated <= 500000{
		taxCalculated = 0.10 * (taxCalculated - 150000)
	}else if taxCalculated <= 1000000{
		taxCalculated = 0.15 * (taxCalculated - 500000)
	}else if taxCalculated <= 2000000{
		taxCalculated = 0.20 * (taxCalculated - 1000000)
	}else if taxCalculated > 2000000{
		taxCalculated =  0.35 * (taxCalculated - 2000000)
	}else {
		return nil, nil
	}

	// wht
	taxCalculated = taxCalculated - request.Wht

	if taxCalculated < 0 {
		taxCalculated = 0
	}

	return &TaxCalculationResponse{Tax: taxCalculated}, nil
} 
