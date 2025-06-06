package comunidaddtos

import (
	"errors"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/commons"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type RequestComunidad struct {
	ID              uint    `json:"id"`
	Nombre          string  `json:"nombre"`
	Descripcion     string  `json:"descripcion"`
	CodigoAcceso    string  `json:"codigo_acceso"`
	Habilitada      *bool   `json:"habilitada"`
	FotoPerfil      string  `json:"foto_perfil"`
	UsuariosID      int64   `json:"usuarios_id"`
	LocalidadID     int64   `json:"localidad_id"`
	TipoComunidadId int64   `json:"tipo_comunidad_id"`
	Email           string  `json:"email"`
	Telefono        string  `json:"telefono"`
	Cuit            string  `json:"cuit"`
	WebUrl          string  `json:"web_url"`
	Calle           string  `json:"calle"`
	StreetAddress   string  `json:"street_address"`
	NumeroPiso      uint    `json:"numero_piso"`
	Lat             float64 `json:"lat"`
	Lng             float64 `json:"lng"`
	Size            int64   `json:"size"`
	Number          int64   `json:"number"`
}

func (r *RequestComunidad) Validate() error {
	if commons.StringIsEmpty(r.Nombre) {
		return errors.New("Por favor, ingresa el nombre de la comunidad.")
	}

	if commons.StringIsEmpty(r.Descripcion) {
		return errors.New("Por favor, proporciona una descripción para la comunidad.")
	}

	if commons.StringIsEmpty(r.Email) {
		return errors.New("Por favor, indica un correo electrónico de contacto.")
	}

	if commons.StringIsEmpty(r.Telefono) {
		return errors.New("Por favor, escribe un número de teléfono válido.")
	}

	if commons.StringIsEmpty(r.Cuit) {
		return errors.New("Por favor, ingresa el CUIT de la comunidad.")
	}

	return nil
}

func (r *RequestComunidad) ToEntity() *entities.Comunidad {
	comunidad := &entities.Comunidad{
		Nombre:          r.Nombre,
		Descripcion:     r.Descripcion,
		CodigoAcceso:    r.CodigoAcceso,
		FotoPerfil:      r.FotoPerfil,
		Email:           r.Email,
		Telefono:        r.Telefono,
		Cuit:            r.Cuit,
		WebUrl:          r.WebUrl,
		StreetAddress:   r.StreetAddress,
		NumeroPiso:      r.NumeroPiso,
		Lat:             r.Lat,
		Lng:             r.Lng,
		TipoComunidadId: r.TipoComunidadId,
		LocalidadesId:   r.LocalidadID,
	}
	comunidad.Habilitada = true
	return comunidad
}
