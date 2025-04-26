package util

import (
	"math"
)

type UtilService interface {
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

func (s *utilService) ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
