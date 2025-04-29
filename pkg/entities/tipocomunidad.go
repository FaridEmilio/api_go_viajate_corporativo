package entities

import "gorm.io/gorm"

type TipoComunidad struct {
	gorm.Model
	Tipo   string `json:"tipo"`
	Activo bool   `json:"activo"`
}

func (TipoComunidad) TableName() string {
	return "tipo_comunidad"
}
