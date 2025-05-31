package entities

import "gorm.io/gorm"

type Trayecto struct {
	gorm.Model
	Alias         string
	Descripcion   string
	Precio        int
	Activo        bool
	OnlyStudents  bool
	OnlyWomen     bool
	VehiculosID   uint
	ComunidadesID uint
	FrecuenciasID uint
	Comunidad     *Comunidad    `gorm:"foreignKey:ComunidadesID"`
	Frecuencia    Frecuencia    `gorm:"foreignKey:FrecuenciasID"`
	Vehiculo      Vehiculo      `gorm:"foreignKey:VehiculosID"`
	Recurrencias  []Recurrencia `gorm:"foreignKey:TrayectosID"`
	Stops         []Stop        `gorm:"foreignKey:TrayectosID"`
}
