package administraciondtos

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type ResponseSedes struct {
	Sedes []ResponseSede `json:"sedes"`
}

type ResponseSede struct {
	ID        uint                          `json:"id"`
	Label     string                        `json:"label"`
	IsCentral bool                          `json:"is_central"`
	Address   comunidaddtos.AddressResponse `json:"address"`
}

func (r *ResponseSedes) FromEntities(sedes []entities.Sede) {
	for _, s := range sedes {
		var sedeDTO ResponseSede
		sedeDTO.FromEntity(s)
		r.Sedes = append(r.Sedes, sedeDTO)
	}
}

func (r *ResponseSede) FromEntity(entity entities.Sede) {
	r.ID = entity.ID
	r.Label = entity.Label
	r.IsCentral = entity.IsCentral

	var addressResp comunidaddtos.AddressResponse
	addressResp.ToAddressResponse(entity.Address)
	r.Address = addressResp
}
