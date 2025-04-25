package entities

import (
	"gorm.io/gorm"
)

type Solicitud struct {
	gorm.Model
	ViajesID   uint    `json:"viajes_id"`
	UsuariosID uint    `json:"usuarios_id"`
	EstadosID  uint    `json:"estados_id"`
	Notificado bool    `json:"notificado"`
	Viaje      Viaje   `json:"viaje" gorm:"foreignKey:ViajesID"`
	Estado     Estado  `json:"estado" gorm:"foreignKey:EstadosID"`
	Usuario    Usuario `json:"usuario" gorm:"foreignKey:UsuariosID"`
}

// TableName sobreescribe el nombre de la tabla
func (Solicitud) TableName() string {
	return "solicitudes"
}
