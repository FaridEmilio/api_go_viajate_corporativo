package authdtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type ResponseUsuario struct {
	ID                   uint    `json:"id"`
	Nombre               string  `json:"nombre"`
	Email                string  `json:"email"`
	EmailVerified        bool    `json:"email_verified"`
	Apellido             string  `json:"apellido"`
	Numero               string  `json:"numero"`
	Genero               string  `json:"genero"`
	FechaNacimiento      string  `json:"fecha_nacimiento"`
	CalificacionPromedio float64 `json:"calificacion_promedio"`
	FotoPerfil           string  `json:"foto_perfil"`
	Activo               bool    `json:"activo"`
	TotalConductor       int64   `json:"total_conductor"`
	TotalPasajero        int64   `json:"total_pasajero"`
	//Comunidades          []comunidaddtos.ResponseComunidad `json:"comunidades"`
	//Reseñas              []ResponseReseña                  `json:"reseñas"`
}

func (r *ResponseUsuario) FromEntity(entity entities.Usuario) {
	r.ID = entity.ID
	r.Nombre = entity.Nombre
	r.Apellido = entity.Apellido
	r.Email = entity.Email
	r.EmailVerified = entity.EmailVerified
	r.Numero = entity.Telefono
	r.FechaNacimiento = entity.FechaNacimiento.Format("2006-01-02")
	r.Genero = entity.Genero
	r.CalificacionPromedio = entity.CalificacionPromedio
	r.FotoPerfil = entity.FotoPerfil
	r.Activo = entity.Activo
}
