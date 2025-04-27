package rutadtos

import (
	"errors"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/commons"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type RequestRuta struct {
	ID               uint    `json:"id"`
	Origen           string  `json:"origen"`
	LatOrigen        float64 `json:"lat_origen"`
	LngOrigen        float64 `json:"lng_origen"`
	CiudadOrigen     string  `json:"ciudad_origen"`
	ProvinciaOrigen  string  `json:"provincia_origen"`
	PaisOrigen       string  `json:"pais_origen"`
	Destino          string  `json:"destino"`
	LatDestino       float64 `json:"lat_destino"`
	LngDestino       float64 `json:"lng_destino"`
	CiudadDestino    string  `json:"ciudad_destino"`
	ProvinciaDestino string  `json:"provincia_destino"`
	PaisDestino      string  `json:"pais_destino"`
	Distancia        float64 `json:"distancia"`
	TiempoEstimado   string  `json:"tiempo_estimado"`
}

func (r *RequestRuta) ToEntity() entities.Ruta {
	return entities.Ruta{
		Origen:           r.Origen,
		Destino:          r.Destino,
		LatOrigen:        r.LatOrigen,
		LngOrigen:        r.LngOrigen,
		LatDestino:       r.LatDestino,
		LngDestino:       r.LngDestino,
		CiudadOrigen:     r.CiudadOrigen,
		ProvinciaOrigen:  r.ProvinciaOrigen,
		PaisOrigen:       r.PaisOrigen,
		CiudadDestino:    r.CiudadDestino,
		ProvinciaDestino: r.ProvinciaDestino,
		PaisDestino:      r.PaisDestino,
		Distancia:        r.Distancia,
		TiempoEstimado:   r.TiempoEstimado,
	}
}

func (r *RequestRuta) Validate() error {
	if commons.StringIsEmpty(r.Origen) {
		return errors.New("El origen es obligatorio")
	}
	if commons.StringIsEmpty(r.Destino) {
		return errors.New("El Destino es obligatorio")
	}

	// Validar que las coordenadas sean números válidos
	if r.LatOrigen == 0 || r.LngOrigen == 0 {
		return errors.New("Las coordenadas de origen son inválidas")
	}
	if r.LatDestino == 0 || r.LngDestino == 0 {
		return errors.New("Las coordenadas de destino son inválidas")
	}

	// Validar que la distancia sea un número positivo
	if r.Distancia <= 0 {
		return errors.New("La distancia es inválida")
	}

	// Validar que el tiempo estimado no esté vacío
	// if commons.StringIsEmpty(r.TiempoEstimado) {
	// 	return errors.New("El campo 'tiempo_estimado' es obligatorio")
	// }

	// // Validar que el campo 'tiempo_estimado' tenga un formato adecuado, ejemplo: "2h 30m"
	// re := regexp.MustCompile(`^\d+h \d+m$`)
	// if !re.MatchString(r.TiempoEstimado) {
	// 	return errors.New("El formato de 'tiempo_estimado' es inválido, debe ser como '2h 30m'")
	// }

	return nil
}
