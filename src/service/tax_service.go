package service

// import "errors"


func TaxCalculation(request TaxCalculationRequest) (*TaxCalculationResponse, error) {
	
	if request.TotalIncome <= 150000 {
		return &TaxCalculationResponse{Tax: 0}, nil
	}else if request.TotalIncome <= 500000{
		return &TaxCalculationResponse{Tax: 0.10 * (request.TotalIncome - 150000 - 60000)}, nil
	}else if request.TotalIncome <= 1000000{
		return &TaxCalculationResponse{Tax: 0.15 * (request.TotalIncome - 500000 - 60000)}, nil
	}else if request.TotalIncome <= 2000000{
		return &TaxCalculationResponse{Tax: 0.20 * (request.TotalIncome - 1000000 - 60000)}, nil
	}else if request.TotalIncome > 2000000{
		return &TaxCalculationResponse{Tax: 0.35 * (request.TotalIncome - 2000000 - 60000)}, nil
	}

	return nil, nil
} 
