package comunidaddtos

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/authdtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type ResponseTrayectos struct {
	Trayectos []TrayectoResponse `json:"trayectos"`
	Meta      dtos.Meta          `json:"meta"`
}
type TrayectoResponse struct {
	ID            uint                     `json:"id"`
	ComunidadesID uint                     `json:"comunidades_id"`
	Frecuencia    string                   `json:"frecuencias_id"`
	Alias         string                   `json:"alias"`
	Descripcion   string                   `json:"descripcion"`
	Precio        int                      `json:"precio"`
	Activo        bool                     `json:"activo"`
	OnlyStudents  bool                     `json:"only_students"`
	OnlyWomen     bool                     `json:"only_women"`
	Usuario       authdtos.ResponseUsuario `json:"usuario"`
	Vehiculo      ResponseVehiculo         `json:"vehiculo"`
	Recurrencias  []RecurrenciaResponse    `json:"recurrencias"`
	Stops         []StopResponse           `json:"stops"`
}

func (r *ResponseTrayectos) ToTrayectosResponse(entities []entities.Trayecto) {
	r.Trayectos = make([]TrayectoResponse, len(entities))
	for i, entity := range entities {
		r.Trayectos[i].ToTrayectoResponse(entity)
	}
}

func (r *TrayectoResponse) ToTrayectoResponse(entity entities.Trayecto) {
	r.ID = entity.ID
	r.ComunidadesID = entity.ComunidadesID
	r.Frecuencia = entity.Frecuencia.Tipo
	r.Alias = entity.Alias
	r.Descripcion = entity.Descripcion
	r.Precio = entity.Precio
	r.Activo = entity.Activo
	r.OnlyStudents = entity.OnlyStudents
	r.OnlyWomen = entity.OnlyWomen
	r.Usuario.FromEntity(entity.Vehiculo.Usuario)
	r.Vehiculo.ToResponseVehiculo(entity.Vehiculo)
	r.ToRecurrenciasResponse(entity.Recurrencias)
	r.ToStopsResponse(entity.Stops)
}

func (r *TrayectoResponse) ToRecurrenciasResponse(entities []entities.Recurrencia) {
	r.Recurrencias = make([]RecurrenciaResponse, len(entities))
	for i, entity := range entities {
		r.Recurrencias[i].ToRecurrenciaResponse(entity)
	}
}

func (r *TrayectoResponse) ToStopsResponse(entities []entities.Stop) {
	r.Stops = make([]StopResponse, len(entities))
	for i, entity := range entities {
		r.Stops[i].ToStopResponse(entity)
	}
}
