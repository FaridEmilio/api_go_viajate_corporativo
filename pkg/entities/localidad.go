package entities

import "gorm.io/gorm"

type Localidad struct {
	gorm.Model
	ProvinciasID uint
	Nombre       string
	Activo       bool
	Provincia    Provincia   `gorm:"foreignKey:ProvinciasID;references:ID"`
}

func (Localidad) TableName() string {
	return "localidades"
}
