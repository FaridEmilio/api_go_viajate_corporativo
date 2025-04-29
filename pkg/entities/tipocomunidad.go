package entities

import "gorm.io/gorm"

type TipoComunidad struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Activo bool   `json:"activo"`
}
