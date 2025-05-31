package comunidaddtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type RequestStop struct {
	Orden   int            `json:"orden"`
	Address RequestAddress `json:"address"`
}

func (r *RequestStop) ToEntity() entities.Stop {
	return entities.Stop{
		Orden:   int64(r.Orden),
		Address: r.Address.ToEntity(),
	}
}
