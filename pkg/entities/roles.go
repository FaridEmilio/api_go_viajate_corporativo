package entities

import "gorm.io/gorm"

type Roles struct {
	gorm.Model
	Rol         string
	Descripcion string
	Usuarios    []Usuario  `gorm:"foreignKey:RolesID"`
	Permisos    []Permisos `gorm:"many2many:roles_has_permisos;"` // Relación de muchos a muchos
}
