package comunidaddtos

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type ResponseTipoComunidad struct {
	Id     uint   `json:"id"`
	Tipo   string `json:"tipo"`
	Activo bool   `json:"activo"`
}

type ResponseTipoComunidades struct {
	TipoComunidades []ResponseTipoComunidad
	Meta            dtos.Meta
}

func (r *ResponseTipoComunidad) FromEntityTipoComunidad(entity entities.TipoComunidad) {
	r.Id = entity.ID
	r.Tipo = entity.Tipo
	r.Activo = entity.Activo
}

func (r *ResponseTipoComunidades) FromEntitiesTipoComunidad(tipocomunidades []entities.TipoComunidad) {
	for _, tipocomunidad := range tipocomunidades {
		var temp ResponseTipoComunidad
		temp.FromEntityTipoComunidad(tipocomunidad)
		r.TipoComunidades = append(r.TipoComunidades, temp)
	}
}
