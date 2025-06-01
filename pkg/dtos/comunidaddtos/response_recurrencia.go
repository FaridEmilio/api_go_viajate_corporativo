package comunidaddtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type RecurrenciaResponse struct {
	ID   uint   `json:"id"`
	Dia  string `json:"dia"`
	Hora string `json:"hora"`
}

func (r *RecurrenciaResponse) ToRecurrenciaResponse(entity entities.Recurrencia) {
	r.ID = entity.ID
	r.Dia = string(entity.Dia)
	r.Hora = entity.Hora
}
