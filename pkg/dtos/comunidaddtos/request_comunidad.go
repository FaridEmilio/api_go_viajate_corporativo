package comunidaddtos

import (
	"errors"
	"strings"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type RequestComunidad struct {
	ID           uint   `json:"id"`
	Nombre       string `json:"nombre"`
	Descripcion  string `json:"descripcion"`
	CodigoAcceso string `json:"codigo_acceso"`
	Habilitada   *bool  `json:"habilitada"`
	FotoPerfil   string `json:"foto_perfil"`
	Size         int64  `json:"size"`
	Number       int64  `json:"number"`
	UsuarioID    int64  `json:"usuario_id"`
}

func (r *RequestComunidad) Validate() error {
	if strings.TrimSpace(r.Nombre) == "" {
		return errors.New("el campo 'nombre' es obligatorio")
	}

	if strings.TrimSpace(r.Descripcion) == "" {
		return errors.New("el campo 'descripcion' es obligatorio")
	}

	return nil
}

func (r *RequestComunidad) ToEntity() *entities.Comunidad {
	comunidad := &entities.Comunidad{
		Nombre:       r.Nombre,
		Descripcion:  r.Descripcion,
		CodigoAcceso: r.CodigoAcceso,
		FotoPerfil:   r.FotoPerfil,
	}

	if r.Habilitada != nil {
		comunidad.Habilitada = *r.Habilitada
	}
	return comunidad
}
