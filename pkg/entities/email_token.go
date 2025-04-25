package entities

import "gorm.io/gorm"

type EmailToken struct {
	gorm.Model
	UsuariosID uint    `json:"usuarios_id"` // ID del usuario que solicita el restablecimiento de contraseña (clave foránea)
	Usuario    Usuario `json:"usuario" gorm:"foreignKey:UsuariosID"`
	Token      string  `json:"token"`
}

func (EmailToken) TableName() string {
	return "email_tokens"
}
