package comunidaddtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type RequestRecurrencia struct {
	Dia  string `json:"dia"`
	Hora string `json:"hora"`
}

func (r *RequestRecurrencia) ToEntity() entities.Recurrencia {
	return entities.Recurrencia{
		Dia:  entities.EnumDia(r.Dia),
		Hora: r.Hora,
	}
}
