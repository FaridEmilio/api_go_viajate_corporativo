package entities

import "gorm.io/gorm"

type Vehiculo struct {
	gorm.Model
	MarcasID   uint   `json:"marcas_id"`
	UsuariosID uint   `json:"usuarios_id"`
	Modelo     string `json:"modelo"`
	Tipo       EnumTipoVehiculo
	Año        int64   `json:"año"`
	Color      string  `json:"color"`
	Patente    string  `json:"patente"`
	Marca      Marca   `json:"marca" gorm:"foreignKey:MarcasID"`
	Usuario    Usuario `json:"usuario" gorm:"foreignKey:UsuariosID"`
}

type EnumTipoVehiculo string

const (
	Moto EnumTipoVehiculo = "Moto"
	Auto EnumTipoVehiculo = "Auto"
)
