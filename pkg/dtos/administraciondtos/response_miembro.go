package administraciondtos

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/commons"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type ResponseMiembros struct {
	Miembros []ResponseMiembro `json:"miembros"`
	Meta     dtos.Meta         `json:"meta"`
}

type ResponseMiembro struct {
	ID                   uint    `json:"id"`
	Activo               bool    `json:"activo"`
	Nombre               string  `json:"nombre"`
	Apellido             string  `json:"apellido"`
	Numero               string  `json:"numero"`
	Genero               string  `json:"genero"`
	FechaNacimiento      string  `json:"fecha_nacimiento"`
	Edad                 int64   `json:"edad"`
	CalificacionPromedio float64 `json:"calificacion_promedio"`
	FotoPerfil           string  `json:"foto_perfil"`
}

func (r *ResponseMiembro) FromEntity(entity entities.Usuario) {
	r.ID = entity.ID
	r.Nombre = entity.Nombre
	r.Apellido = entity.Apellido
	r.Numero = entity.Telefono
	r.FechaNacimiento = entity.FechaNacimiento.Format("2006-01-02")
	r.Genero = entity.Genero
	r.CalificacionPromedio = entity.CalificacionPromedio
	r.FotoPerfil = entity.FotoPerfil
	r.Edad = commons.CalcularEdad(entity.FechaNacimiento)
}

func (r *ResponseMiembros) ToMiembrosResponse(entities []entities.UsuariosHasComunidades) {
	r.Miembros = make([]ResponseMiembro, len(entities))
	for i, entity := range entities {
		r.Miembros[i].FromEntity(entity.Usuario)
		r.Miembros[i].Activo = entity.Activo
	}
}
