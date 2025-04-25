package entities

import (
	"gorm.io/gorm"
)

type Comunidad struct {
	gorm.Model
	Nombre       string
	Descripcion  string
	Usuarios     []Usuario `gorm:"many2many:usuarios_has_comunidades;"` // Relación de muchos a muchos
	CodigoAcceso string
	Habilitada   bool
	FotoPerfil   bool
}

func (Comunidad) TableName() string {
	return "comunidades"
}
