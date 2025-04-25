package entities

import "gorm.io/gorm"

type NotificationToken struct {
	gorm.Model
	UsuariosID uint    `json:"usuarios_id"`
	Token      string  `json:"token"`
	Usuario    Usuario `json:"usuario" gorm:"foreignKey:UsuariosID"`
}
