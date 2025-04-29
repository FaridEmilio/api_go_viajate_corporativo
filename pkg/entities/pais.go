package entities

import "gorm.io/gorm"

type Pais struct {
	gorm.Model
	Nombre     string
	Activo     bool
	Provincias []Provincia `gorm:"foreignKey:PaisesID;references:ID"`
}

func (Pais) TableName() string {
	return "paises"
}
