package comunidaddtos

import (
	"github.com/faridEmilio/api_go_viajate/pkg/commons"
	"github.com/faridEmilio/api_go_viajate/pkg/entities"
)

type ResponseUsuarioComunidad struct {
	ID                   uint    `json:"id"`
	Nombre               string  `json:"nombre"`
	Apellido             string  `json:"apellido"`
	Numero               string  `json:"numero"`
	Genero               string  `json:"genero"`
	FechaNacimiento      string  `json:"fecha_nacimiento"`
	Edad                 int64   `json:"edad"`
	CalificacionPromedio float64 `json:"calificacion_promedio"`
	FotoPerfil           string  `json:"foto_perfil"`
}

func (r *ResponseUsuarioComunidad) FromEntity(entity entities.Usuario) {
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

