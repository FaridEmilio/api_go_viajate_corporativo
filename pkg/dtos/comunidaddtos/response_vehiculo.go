package comunidaddtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type ResponseVehiculos struct {
	Vehiculos []ResponseVehiculo `json:"vehiculos"`
}

type ResponseVehiculo struct {
	ID      uint   `json:"id"`
	Marca   string `json:"marca"`
	Modelo  string `json:"modelo"`
	Año     int64  `json:"año"`
	Color   string `json:"color"`
	Patente string `json:"patente"`
	Tipo    string `json:"tipo"`
}

func (r *ResponseVehiculo) ToResponseVehiculo(entity entities.Vehiculo) {
	r.ID = entity.ID
	r.Marca = entity.Marca.Marca
	r.Modelo = entity.Modelo
	r.Año = entity.Año
	r.Color = entity.Color
	r.Patente = entity.Patente
	r.Tipo = string(entity.Tipo)
}

func (r *ResponseVehiculos) ToResponseVehiculos(entities []entities.Vehiculo) {
	r.Vehiculos = make([]ResponseVehiculo, len(entities))
	for i, entity := range entities {
		r.Vehiculos[i].ToResponseVehiculo(entity)
	}
}
