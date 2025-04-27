package rutadtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type ResponseRuta struct {
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

func (r *ResponseRuta) ToRutaResponse(entity entities.Ruta) {
	r.ID = entity.ID
	r.Origen = entity.Origen
	r.LatOrigen = entity.LatOrigen
	r.LngOrigen = entity.LngOrigen
	r.CiudadOrigen = entity.CiudadOrigen
	r.ProvinciaOrigen = entity.ProvinciaOrigen
	r.PaisOrigen = entity.PaisOrigen
	r.Destino = entity.Destino
	r.LatDestino = entity.LatDestino
	r.LngDestino = entity.LngDestino
	r.CiudadDestino = entity.CiudadDestino
	r.ProvinciaDestino = entity.ProvinciaDestino
	r.PaisDestino = entity.PaisDestino
	r.Distancia = entity.Distancia
	r.TiempoEstimado = entity.TiempoEstimado
}
