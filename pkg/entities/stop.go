package entities

import "gorm.io/gorm"

type Stop struct {
	gorm.Model
	Orden       int64
	TrayectosID uint
	AddressesID uint
	Trayecto    Trayecto `gorm:"foreignKey:TrayectosID"`
	Address     Address  `gorm:"foreignKey:AddressesID"`
}
