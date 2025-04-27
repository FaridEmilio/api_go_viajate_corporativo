package entities

import (
	"gorm.io/gorm"
)

type Comunidad struct {
	gorm.Model
	Nombre       string    `json:"nombre"`
	Descripcion  string    `json:"descripcion"`
	Usuarios     []Usuario `gorm:"many2many:usuarios_has_comunidades;"` // Relación de muchos a muchos
	CodigoAcceso string    `json:"codigo_acceso"`
	Habilitada   bool      `json:"habilitada"`
	FotoPerfil   string    `json:"foto_perfil"`
}

func (Comunidad) TableName() string {
	return "comunidades"
}
