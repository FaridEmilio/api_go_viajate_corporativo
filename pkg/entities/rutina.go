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
	Trayecto    []Trayecto `gorm:"many2many:trayectos_has_rutinas;"` // Relación de muchos a muchos
}
