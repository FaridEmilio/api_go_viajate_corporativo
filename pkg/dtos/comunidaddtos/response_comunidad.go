package comunidaddtos

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type ResponseComunidades struct {
	Comunidades []ResponseComunidad `json:"comunidades"`
	Meta        dtos.Meta           `json:"meta"`
}

type ResponseComunidad struct {
	ID           uint   `json:"id"`
	Nombre       string `json:"nombre"`
	Descripcion  string `json:"descripcion"`
	CodigoAcceso string `json:"codigo_acceso"`
	Habilitada   bool   `json:"habilitada"`
	FotoPerfil   string `json:"foto_perfil"`
}

func (r *ResponseComunidad) FromEntity(entity entities.Comunidad) {
	r.ID = entity.ID
	r.Nombre = entity.Nombre
	r.Descripcion = entity.Descripcion
	r.CodigoAcceso = entity.CodigoAcceso
	r.Habilitada = entity.Habilitada
	r.FotoPerfil = entity.FotoPerfil
}

func (r *ResponseComunidades) FromEntities(comunidades []entities.Comunidad) {
	for _, comunidad := range comunidades {
		var temp ResponseComunidad
		temp.FromEntity(comunidad)
		r.Comunidades = append(r.Comunidades, temp)
	}
}
