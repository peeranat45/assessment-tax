package service


func TaxCalculation(request TaxCalculationRequest) (*TaxCalculationResponse, error) {
	response := TaxCalculationResponse{
		Tax: 29000,
	}
	return &response, nil
} 
