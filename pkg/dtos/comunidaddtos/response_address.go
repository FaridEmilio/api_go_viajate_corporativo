package comunidaddtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type AddressResponse struct {
	ID               uint    `json:"id"`
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

func (r *AddressResponse) ToAddressResponse(entity entities.Address) {
	r.ID = entity.ID
	r.Name = entity.Name
	r.Lat = entity.Lat
	r.Lng = entity.Lng
	r.StreetAddress = entity.StreetAddress
	r.FormattedAddress = entity.FormattedAddress
	r.PostalCode = entity.PostalCode
	r.City = entity.City
	r.Province = entity.Province
	r.Country = entity.Country
	r.URL = entity.Url
}
