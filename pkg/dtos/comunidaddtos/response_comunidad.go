package comunidaddtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type ResponseComunidades struct {
	Comunidades []ResponseComunidad `json:"comunidades"`
}

type ResponseComunidad struct {
	ID          uint   `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

func (r *ResponseComunidad) FromEntity(entity entities.Comunidad) {
	r.ID = entity.ID
	r.Nombre = entity.Nombre
	r.Descripcion = entity.Descripcion
}

func (r *ResponseComunidades) FromEntities(comunidades []entities.Comunidad) {
	for _, comunidad := range comunidades {
		var temp ResponseComunidad
		temp.FromEntity(comunidad)
		r.Comunidades = append(r.Comunidades, temp)
	}
}
