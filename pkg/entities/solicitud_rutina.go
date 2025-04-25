package entities

import "gorm.io/gorm"

type SolicitudRutina struct {
	gorm.Model
	RutinasID  uint    `json:"rutinas_id"`
	UsuariosID uint    `json:"usuarios_id"`
	EstadosID  uint    `json:"estados_id"`
	Notificado bool    `json:"notificado"`
	Rutina     Rutina  `json:"rutina" gorm:"foreignKey:RutinasID"`
	Estado     Estado  `json:"estado" gorm:"foreignKey:EstadosID"`
	Usuario    Usuario `json:"usuario" gorm:"foreignKey:UsuariosID"`
}
