package entities

import (
	"time"

	"gorm.io/gorm"
)

type Viaje struct {
	gorm.Model
	Uuid        string      `json:"uuid"`
	UsuariosID  uint        `json:"usuarios_id"`
	EstadosID   uint        `json:"estados_id"`
	RutasID     uint        `json:"rutas_id"`
	VehiculosID uint        `json:"vehiculos_id"`
	Fecha       time.Time   `json:"fecha"`
	Hora        string      `json:"hora"`
	Asientos    int64       `json:"asientos"`
	Precio      int64       `json:"precio"`
	Mascotas    bool        `json:"mascotas"`
	Equipaje    bool        `json:"equipaje"`
	Descripcion string      `json:"descripcion"`
	Usuario     Usuario     `json:"usuario" gorm:"foreignKey:UsuariosID"`
	Estado      Estado      `json:"estado" gorm:"foreignKey:EstadosID"`
	Vehiculo    Vehiculo    `json:"vehiculo" gorm:"foreignKey:VehiculosID"`
	Solicitudes []Solicitud `json:"solicitudes" gorm:"foreignKey:ViajesID"`
}
