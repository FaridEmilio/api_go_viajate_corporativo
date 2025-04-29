package comunidaddtos

import (
	"errors"
	"strings"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type RequestComunidad struct {
	ID              uint    `json:"id"`
	Nombre          string  `json:"nombre"`
	Descripcion     string  `json:"descripcion"`
	CodigoAcceso    string  `json:"codigo_acceso"`
	Habilitada      *bool   `json:"habilitada"`
	FotoPerfil      string  `json:"foto_perfil"`
	UsuarioID       int64   `json:"usuario_id"`
	LocalidadId     int64   `json:"localidad_id"`
	TipoComunidadId int64   `json:"tipo_comunidad_id"`
	Email           string  `json:"email"`
	Telefono        string  `json:"telefono"`
	Cuit            string  `json:"cuit"`
	WebUrl          string  `json:"web_url"`
	Calle           string  `json:"calle"`
	Altura          int     `json:"altura"`
	NumeroPiso      uint    `json:"numero_piso"`
	Lat             float64 `json:"lat"`
	Lng             float64 `json:"lng"`
	Size            int64   `json:"size"`
	Number          int64   `json:"number"`
}

func (r *RequestComunidad) Validate() error {
	if strings.TrimSpace(r.Nombre) == "" {
		return errors.New("el campo 'nombre' es obligatorio")
	}

	if strings.TrimSpace(r.Descripcion) == "" {
		return errors.New("el campo 'descripcion' es obligatorio")
	}

	if strings.TrimSpace(r.Email) == "" {
		return errors.New("el campo 'email' es obligatorio")
	}

	if strings.TrimSpace(r.Telefono) == "" {
		return errors.New("el campo 'telefono' es obligatorio")
	}

	if strings.TrimSpace(r.Cuit) == "" {
		return errors.New("el campo 'cuit' es obligatorio")
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
		Calle:           r.Calle,
		Altura:          r.Altura,
		NumeroPiso:      r.NumeroPiso,
		Lat:             r.Lat,
		Lng:             r.Lng,
		TipoComunidadId: r.TipoComunidadId,
		LocalidadesId:   r.LocalidadId,
	}
	return comunidad
}
