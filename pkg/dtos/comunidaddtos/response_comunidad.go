package comunidaddtos

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type ResponseComunidades struct {
	Comunidades []ResponseComunidad `json:"comunidades"`
	Meta        dtos.Meta           `json:"meta"`
}

type ResponseComunidad struct {
	ID            uint    `json:"id"`
	Nombre        string  `json:"nombre"`
	Descripcion   string  `json:"descripcion"`
	CodigoAcceso  string  `json:"codigo_acceso"`
	Habilitada    bool    `json:"habilitada"`
	FotoPerfil    string  `json:"foto_perfil"`
	Localidad     string  `json:"localidad"`
	Provincia     string  `json:"provincia"`
	Pais          string  `json:"pais"`
	TipoComunidad string  `json:"tipo_comunidad"`
	Email         string  `json:"email"`
	Telefono      string  `json:"telefono"`
	Cuit          string  `json:"cuit"`
	WebUrl        string  `json:"web_url"`
	Calle         string  `json:"calle"`
	Altura        int     `json:"altura"`
	NumeroPiso    uint    `json:"numero_piso"`
	Lat           float64 `json:"lat"`
	Lng           float64 `json:"lng"`
}

func (r *ResponseComunidad) FromEntity(entity entities.Comunidad) {
	r.ID = entity.ID
	r.Nombre = entity.Nombre
	r.Descripcion = entity.Descripcion
	r.CodigoAcceso = entity.CodigoAcceso
	r.Habilitada = entity.Habilitada
	r.FotoPerfil = entity.FotoPerfil
	r.Localidad = entity.Localidad.Nombre
	r.Provincia = entity.Localidad.Provincia.Nombre
	r.Pais = entity.Localidad.Provincia.Pais.Nombre
	r.TipoComunidad = entity.TipoComunidad.Tipo
	r.Email = entity.Email
	r.Telefono = entity.Telefono
	r.Cuit = entity.Cuit
	r.WebUrl = entity.WebUrl
	r.Calle = entity.Calle
	r.Altura = entity.Altura
	r.NumeroPiso = entity.NumeroPiso
	r.Lat = entity.Lat
	r.Lng = entity.Lng
}

func (r *ResponseComunidades) FromEntities(comunidades []entities.Comunidad) {
	for _, comunidad := range comunidades {
		var temp ResponseComunidad
		temp.FromEntity(comunidad)
		r.Comunidades = append(r.Comunidades, temp)
	}
}
