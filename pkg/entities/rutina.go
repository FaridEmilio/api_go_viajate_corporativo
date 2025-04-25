package entities

import (
	"time"

	"gorm.io/gorm"
)

type Rutina struct {
	gorm.Model
	TrayectosID uint      `json:"trayectos_id"`
	Dia         string    `json:"día"`
	Hora        string    `json:"hora"`
	Fecha       time.Time `json:"fecha"`
	Trayecto    *Trayecto `json:"trayecto" gorm:"foreignKey:TrayectosID"`
}
