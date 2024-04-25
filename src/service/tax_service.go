package service

// import "errors"


func TaxCalculation(request TaxCalculationRequest) (*TaxCalculationResponse, error) {

	taxCalculated := 0.0
	
	if request.TotalIncome <= 150000 {
		taxCalculated = 0
	}else if request.TotalIncome <= 500000{
		taxCalculated = 0.10 * (request.TotalIncome - 150000 - 60000)
	}else if request.TotalIncome <= 1000000{
		taxCalculated = 0.15 * (request.TotalIncome - 500000 - 60000)
	}else if request.TotalIncome <= 2000000{
		taxCalculated = 0.20 * (request.TotalIncome - 1000000 - 60000)
	}else if request.TotalIncome > 2000000{
		taxCalculated =  0.35 * (request.TotalIncome - 2000000 - 60000)
	}else {
		return nil, nil
	}

	return &TaxCalculationResponse{Tax: taxCalculated}, nil
} 
