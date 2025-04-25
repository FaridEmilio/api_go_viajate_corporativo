package entities

import "gorm.io/gorm"

type Historial struct {
	gorm.Model
	UsuariosID    uint    `json:"usuarios_id"`
	Usuario       Usuario `json:"usuario" gorm:"foreignKey:UsuariosID"`
	Funcionalidad string  `json:"funcionalidad"`
	Descripcion   string  `json:"descripcion"`
	Contacto      string  `json:"contacto"`
}

// TableName sobreescribe el nombre de la tabla
func (Historial) TableName() string {
	return "historial"
}
