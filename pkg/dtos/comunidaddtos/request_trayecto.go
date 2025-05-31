package comunidaddtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type RequestTrayecto struct {
	ComunidadesID int                  `json:"comunidades_id"`
	VehiculosID   int                  `json:"vehiculos_id"`
	FrecuenciasID int                  `json:"frecuencias_id"`
	Alias         string               `json:"alias"`
	Descripcion   string               `json:"descripcion"`
	Precio        int                  `json:"precio"`
	OnlyStudents  bool                 `json:"only_students"`
	OnlyWomen     bool                 `json:"only_women"`
	Recurrencias  []RequestRecurrencia `json:"recurrencias"`
	Stops         []RequestStop        `json:"stops"`
}

func (r *RequestTrayecto) ToEntity() entities.Trayecto {
	return entities.Trayecto{
		ComunidadesID: uint(r.ComunidadesID),
		VehiculosID:   uint(r.VehiculosID),
		FrecuenciasID: uint(r.FrecuenciasID),
		Alias:         r.Alias,
		Descripcion:   r.Descripcion,
		Precio:        r.Precio,
		OnlyStudents:  r.OnlyStudents,
		OnlyWomen:     r.OnlyWomen,
		Activo:        true,
		Recurrencias:  r.ToRecurrenciasEntities(),
		Stops:         r.ToStopsEntities(),
	}
}

func (r *RequestTrayecto) ToRecurrenciasEntities() []entities.Recurrencia {
	recurrencias := make([]entities.Recurrencia, len(r.Recurrencias))
	for i := range r.Recurrencias {
		recurrencias[i] = r.Recurrencias[i].ToEntity()
	}
	return recurrencias
}

func (r *RequestTrayecto) ToStopsEntities() []entities.Stop {
	stops := make([]entities.Stop, len(r.Stops))
	for i := range r.Stops {
		stops[i] = r.Stops[i].ToEntity()
	}
	return stops
}
