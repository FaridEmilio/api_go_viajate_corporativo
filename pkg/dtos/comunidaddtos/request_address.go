package comunidaddtos

import (
	"errors"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type RequestAddress struct {
	Name             string  `json:"name"`
	Lat              float64 `json:"lat"`
	Lng              float64 `json:"lng"`
	StreetAddress    string  `json:"street_address"`
	FormattedAddress string  `json:"formatted_address"`
	PostalCode       string  `json:"postal_code"`
	City             string  `json:"city"`
	Province         string  `json:"province"`
	Country          string  `json:"country"`
	URL              string  `json:"url,omitempty"`
}

func (r *RequestAddress) ToEntity() entities.Address {
	return entities.Address{
		Name:             r.Name,
		Lat:              r.Lat,
		Lng:              r.Lng,
		StreetAddress:    r.StreetAddress,
		FormattedAddress: r.FormattedAddress,
		PostalCode:       r.PostalCode,
		City:             r.City,
		Province:         r.Province,
		Country:          r.Country,
		Url:              r.URL,
	}
}

func (r *RequestAddress) Validate() error {
	if r.StreetAddress == "" {
		return errors.New("el campo 'street_address' es obligatorio")
	}
	if r.City == "" {
		return errors.New("el campo 'city' es obligatorio")
	}
	if r.Country == "" {
		return errors.New("el campo 'country' es obligatorio")
	}
	return nil
}
