package entities

import "gorm.io/gorm"

type UsuarioHasReseña struct {
	gorm.Model
	ReseñasID  uint    `json:"reseñas_id"`
	UsuariosID uint    `json:"usuarios_id"`
	Usuario    Usuario `json:"usuario" gorm:"foreignKey:UsuariosID"`
	Reseña     Reseña  `json:"reseña" gorm:"foreignKey:ReseñasID"`
}
