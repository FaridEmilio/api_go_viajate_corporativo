package entities

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Name             string
	Lat              float64
	Lng              float64
	StreetAddress    string
	FormattedAddress string
	PostalCode       string
	City             string
	Province         string
	Country          string
	Url              string
}

func (Address) TableName() string {
	return "addresses"
}
