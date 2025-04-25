package entities

import "gorm.io/gorm"

type Ruta struct {
	gorm.Model
	// Origen Data
	Origen          string  `json:"origen"`     // Nombre del origen
	LatOrigen       float64 `json:"lat_origen"` // Latitud del origen
	LngOrigen       float64 `json:"lng_origen"` // Longitud del origen
	CiudadOrigen    string  `json:"ciudad_origen"`
	ProvinciaOrigen string  `json:"provincia_origen"`
	PaisOrigen      string  `json:"pais_origen"`
	// Destino Data
	Destino          string  `json:"destino"`     // Nombre del destino
	LatDestino       float64 `json:"lat_destino"` // Latitud del destino
	LngDestino       float64 `json:"lng_destino"` // Longitud del destino
	CiudadDestino    string  `json:"ciudad_destino"`
	ProvinciaDestino string  `json:"provincia_destino"`
	PaisDestino      string  `json:"pais_destino"`
	// EXtra Data
	Distancia      float64 `json:"distancia"`       // Distancia entre el origen y destino (en km o millas)
	TiempoEstimado string  `json:"tiempo_estimado"` // Tiempo estimado para el trayecto (en formato "2h 30m" o en minutos)
}

// TableName sobreescribe el nombre de la tabla
func (Ruta) TableName() string {
	return "rutas"
}
