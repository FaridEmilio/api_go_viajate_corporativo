package entities

import "gorm.io/gorm"

type Recurrencia struct {
	gorm.Model
	TrayectosID uint
	Dia         EnumDia
	Hora        string
}

type EnumDia string

const (
	Lunes     EnumDia = "Lunes"
	Martes    EnumDia = "Martes"
	Miercoles EnumDia = "Miércoles"
	Jueves    EnumDia = "Jueves"
	Viernes   EnumDia = "Viernes"
	Sabado    EnumDia = "Sábado"
	Domingo   EnumDia = "Domingo"
)
