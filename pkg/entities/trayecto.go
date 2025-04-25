package entities

import (
	"gorm.io/gorm"
)

type Trayecto struct {
	gorm.Model
	ComunidadesID uint            `json:"comunidades_id"`
	UsuariosID    uint            `json:"usuarios_id"`
	RutasID       uint            `json:"rutas_id"`
	EstadosID     uint            `json:"estados_id"`
	Precio        uint            `json:"precio"`
	Uuid          string          `json:"uuid"`
	Descripcion   string          `json:"descripcion"`
	Detalles      *TrayectoDetalle `json:"trayecto_detalles" gorm:"foreignKey:TrayectosID"`
	Estado        Estado          `json:"estado" gorm:"foreignKey:EstadosID"`
	Ruta          Ruta            `json:"ruta" gorm:"foreignKey:RutasID"`
	Comunidad     Comunidad       `json:"comunidad" gorm:"foreignKey:ComunidadesID"`
	Usuario       Usuario         `json:"usuario" gorm:"foreignKey:UsuariosID"`
	Rutinas       []*Rutina        `json:"rutinas" gorm:"foreignKey:TrayectosID"`
}
