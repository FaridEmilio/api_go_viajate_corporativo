package comunidaddtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type ResponseMarcas struct {
	Marcas []ResponseMarca `json:"marcas"`
}
type ResponseMarca struct {
	ID    uint   `json:"id"`
	Marca string `json:"marca"`
}

func (r *ResponseMarca) ToResponseMarca(entity entities.Marca) {
	r.ID = entity.ID
	r.Marca = entity.Marca
}

func (r *ResponseMarcas) ToResponseMarcas(entities []entities.Marca) {
	r.Marcas = make([]ResponseMarca, len(entities))
	for i, entity := range entities {
		r.Marcas[i].ToResponseMarca(entity)
	}
}
