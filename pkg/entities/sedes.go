package entities

import "gorm.io/gorm"

type Sede struct {
	gorm.Model
	ID            uint
	ComunidadesID uint
	AddressesID   uint
	Label         string
	Active        bool
	IsCentral     bool
	Address       Address `gorm:"foreignKey:AddressesID"` // para preload
}
