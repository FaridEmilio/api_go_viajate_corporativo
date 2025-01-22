package util

import (
	"fmt"
	"math"

	"github.com/faridEmilio/api_go_gym_manager/internal/logs"
	"github.com/faridEmilio/api_go_gym_manager/pkg/entities"
)

type UtilService interface {
	CreateNotificacionService(notificacion entities.Notificacione) (erro error)
	CreateLogService(log entities.Log) (erro error)
	LogError(erro string, funcionalidad string)
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

func (r *utilService) CreateNotificacionService(notificacion entities.Notificacione) (erro error) {
	return r.repository.CreateNotificacion(notificacion)

}

func (r *utilService) CreateLogService(log entities.Log) (erro error) {
	return r.repository.CreateLog(log)
}

func (r *utilService) LogError(erro string, funcionalidad string) {

	log := entities.Log{
		Tipo:          entities.Error,
		Mensaje:       erro,
		Funcionalidad: funcionalidad,
	}

	err := r.CreateLogService(log)

	if err != nil {
		mensaje := fmt.Sprintf("Crear Log: %s. %s", err.Error(), erro)
		logs.Error(mensaje)
	}
}

func (s *utilService) ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
