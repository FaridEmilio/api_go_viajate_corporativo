package entities

import "gorm.io/gorm"

type TrayectoDetalle struct {
	gorm.Model
	TrayectosID       uint      `json:"trayectos_id"`
	FrecuenciaSemanal bool      `json:"frecuencia_semanal"`
	SoloEstudiantes   bool      `json:"solo_estudiantes"`
	SoloMujeres       bool      `json:"solo_mujeres"`
	Mascotas          bool      `json:"mascotas"`
	Asientos          uint      `json:"asientos"`
	Trayecto          *Trayecto `json:"trayecto" gorm:"foreignKey:TrayectosID"`
}
