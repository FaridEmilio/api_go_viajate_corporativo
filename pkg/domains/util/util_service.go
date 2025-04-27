package util

import (
	"math"
)

type UtilService interface {
	//RunEndpoint(method, endpoint string, headers map[string]string, body interface{}, queryParams map[string]string, logRequest bool, response interface{}) error

	//Redondeo
	ToFixed(num float64, precision int) float64
}

func NewUtilService(r UtilRepository) UtilService {
	service := utilService{
		repository: r,
	}
	return &service
}

type utilService struct {
	repository UtilRepository
}

// func (s *utilService) RunEndpoint(method string, endpoint string, headers map[string]string, body interface{}, queryParams map[string]string, logRequest bool, response interface{}) error {

// }

func (s *utilService) ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
