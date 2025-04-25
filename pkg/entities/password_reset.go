package entities

import "gorm.io/gorm"

type PasswordReset struct {
	gorm.Model
	UsuariosID uint    `json:"usuarios_id"` // ID del usuario que solicita el restablecimiento de contraseña (clave foránea)
	Usuario    Usuario `json:"usuario" gorm:"foreignKey:UsuariosID"`
	Token      string  `json:"token"` // Token único para el restablecimiento de contraseña
}
