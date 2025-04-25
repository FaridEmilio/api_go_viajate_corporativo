package entities

import (
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Uuid                 string                `json:"uuid"`
	Nombre               string                `json:"nombre"`
	Apellido             string                `json:"apellido"`
	Email                string                `gorm:"unique;not null" json:"email"`
	EmailVerified        bool                  `json:"email_verified"`
	Activo               bool                  `json:"activo"`
	FechaNacimiento      time.Time             `json:"fecha_nacimiento"`
	Telefono             string                `json:"telefono"`
	Genero               string                `json:"genero"`
	CalificacionPromedio float64               `json:"calificacion_promedio"`
	Terminos             bool                  `json:"terminos"`
	FotoPerfil           string                `json:"foto_perfil"`
	Contraseña           []byte                `json:"contraseña"`
	Viajes               []Viaje               `json:"viajes" gorm:"foreignKey:UsuariosID"`
	Solicitudes          []Solicitud           `json:"solicitudes" gorm:"foreignKey:UsuariosID"`
	RolesComunidad       []UsuarioRolComunidad `json:"roles_comunidades" gorm:"foreignKey:UsuariosID"`
	UsuarioHasReseña     []UsuarioHasReseña    `json:"reseñas_recibidas" gorm:"foreignKey:UsuariosID"`
	Reseña               []Reseña              `json:"reseñas_creadas" gorm:"foreignKey:UsuariosID"`
	NotificationTokens   []NotificationToken   `json:"notification_tokens" gorm:"foreignKey:UsuariosID"`
	Vehiculos            []Vehiculo            `json:"vehiculos" gorm:"foreignKey:UsuariosID"`
}
