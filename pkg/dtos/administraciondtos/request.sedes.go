package administraciondtos

import (
	"errors"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type RequestCreateSede struct {
	Id          uint                         `json:"id"`
	Label       string                       `json:"label"`
	IsCentral   bool                         `json:"is_central"`
	ComunidadID uint                         `json:"comunidad_id"`
	Active      bool                         `json:"active"`
	Address     comunidaddtos.RequestAddress `json:"address"`
}

func (r *RequestCreateSede) Validate() error {
	if r.Label == "" {
		return errors.New("el campo 'label' es obligatorio")
	}
	if r.ComunidadID == 0 {
		return errors.New("el campo 'comunidad_id' es obligatorio")
	}
	if err := r.Address.Validate(); err != nil {
		return err
	}
	return nil
}

func (r *RequestCreateSede) ToEntity() (entities.Sede, entities.Address) {
	address := r.Address.ToEntity()
	sede := entities.Sede{
		Label:         r.Label,
		IsCentral:     r.IsCentral,
		ComunidadesID: r.ComunidadID,
		Active:        true,
	}
	return sede, address
}
