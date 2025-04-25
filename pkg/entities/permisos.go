package entities

import "gorm.io/gorm"

type Permisos struct {
	gorm.Model
	Permiso     string
	Descripcion string
	Roles       []Roles `gorm:"many2many:roles_has_permisos;"` // Relación de muchos a muchos
}
