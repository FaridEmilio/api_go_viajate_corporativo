package administraciondtos

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"

type ResponsePaises struct {
	Paises []ResponsePais `json:"paises"`
}

type ResponsePais struct {
	Nombre     string              `json:"nombre"`
	Provincias []ResponseProvincia `json:"provincias"`
}

type ResponseProvincia struct {
	Nombre      string              `json:"nombre"`
	Localidades []ResponseLocalidad `json:"localidades"`
}

type ResponseLocalidad struct {
	ID     uint   `json:"id"`
	Nombre string `json:"nombre"`
}

func (r *ResponseLocalidad) FromEntity(entity entities.Localidad) {
	r.ID = entity.ID
	r.Nombre = entity.Nombre
}

func (r *ResponseProvincia) FromEntity(entity entities.Provincia) {
	r.Nombre = entity.Nombre
	r.Localidades = make([]ResponseLocalidad, len(entity.Localidades))
	for i := range entity.Localidades {
		r.Localidades[i].FromEntity(entity.Localidades[i])
	}
}

func (r *ResponsePais) FromEntity(entity entities.Pais) {
	r.Nombre = entity.Nombre
	r.Provincias = make([]ResponseProvincia, len(entity.Provincias))
	for i := range entity.Provincias {
		r.Provincias[i].FromEntity(entity.Provincias[i])
	}
}

func (r *ResponsePaises) FromEntities(entities []entities.Pais) {
	r.Paises = make([]ResponsePais, len(entities))
	for i := range entities {
		r.Paises[i].FromEntity(entities[i])
	}
}
