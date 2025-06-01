package comunidaddtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type StopResponse struct {
	ID      uint            `json:"id"`
	Orden   int64           `json:"orden"`
	Address AddressResponse `json:"address"`
}

func (r *StopResponse) ToStopResponse(entity entities.Stop) {
	r.ID = entity.ID
	r.Orden = entity.Orden
	r.Address.ToAddressResponse(entity.Address)
}
