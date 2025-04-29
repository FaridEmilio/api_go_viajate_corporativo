package entities

import "gorm.io/gorm"

type Provincia struct {
	gorm.Model
	PaisesID    uint
	Nombre      string
	Activo      bool
	Pais        Pais        `gorm:"foreignKey:PaisesID;references:ID"`
	Localidades []Localidad `gorm:"foreignKey:ProvinciasID;references:ID"`
}

func (Provincia) TableName() string {
	return "provincias"
}
