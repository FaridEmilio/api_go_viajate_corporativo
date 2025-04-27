package usuariosdtos

import (
	"errors"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/commons"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/viajatedtos"
)

type RequestUpdateUser struct {
	Nombre          string                 `json:"nombre"`
	Apellido        string                 `json:"apellido"`
	FechaNacimiento string                 `json:"fecha_nacimiento"`
	Genero          viajatedtos.EnumGenero `json:"genero"`
}

// Valida todos los campos de RequestUpdateUser
func (req *RequestUpdateUser) Validate() error {
	// Nombre
	if commons.StringIsEmpty(req.Nombre) {
		return errors.New("El nombre es obligatorio")
	}
	if err := commons.IsNameValid(req.Nombre); err != nil {
		return err
	}
	// Apellido
	if commons.StringIsEmpty(req.Apellido) {
		return errors.New("El apellido es obligatorio")
	}
	if err := commons.IsNameValid(req.Apellido); err != nil {
		return err
	}
	// Fecha de Nacimiento
	if commons.StringIsEmpty(req.FechaNacimiento) {
		return errors.New("La fecha de nacimiento es obligatoria")
	}
	// Genero
	if err := req.Genero.IsGeneroValid(); err != nil {
		return err
	}
	return nil
}
