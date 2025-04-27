package entities

import (
	"gorm.io/gorm"
)

type Reseña struct {
	gorm.Model
	UsuariosID       uint              `json:"usuarios_id"`
	RolViaje         string            `json:"rol_viaje"`
	Valoracion       int64             `json:"valoracion"` // Se envian valores del 1 al 5
	Reseña           string            `json:"reseña"`
	Usuario          Usuario           `json:"usuario" gorm:"foreignKey:UsuariosID"` // usuario creador de la reseña
}
