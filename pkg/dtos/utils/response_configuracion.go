package utilsdtos

import (
	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/pkg/dtos"
	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/pkg/entities"
)

type ResponseConfiguraciones struct {
	Data []ResponseConfiguracion `json:"data"`
	Meta dtos.Meta               `json:"meta"`
}

type ResponseConfiguracion struct {
	Id          uint
	Nombre      string
	Descripcion string
	Valor       string
}

func (r *ResponseConfiguracion) FromEntity(c entities.Configuracione) {
	r.Id = c.ID
	r.Nombre = c.Nombre
	r.Descripcion = c.Descripcion
	r.Valor = c.Valor
}
