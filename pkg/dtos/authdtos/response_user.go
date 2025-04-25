package authdtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type ResponseUser struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Activo      bool   `json:"activo"`
	ProvincesID uint   `json:"provinces_id"`
}

func (r *ResponseUser) FromEntity(entity entities.Users) {
	r.ID = entity.ID
	r.Name = entity.Name
	r.Email = entity.Email
	r.Activo = entity.Activo
	r.ProvincesID = entity.ProvincesID
}
