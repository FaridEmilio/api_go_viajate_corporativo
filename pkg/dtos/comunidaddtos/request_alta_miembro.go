package comunidaddtos

import (
	"errors"
	"regexp"
	"strings"
)

type RequestMiembro struct {
	Codigo      string `json:"codigo"`
	UsuariosID  uint   `json:"usuarios_id"`
	ComunidadID uint   `json:"comunidad_id"`
	Activo      *bool  `json:"activo"`
}

func (r *RequestMiembro) IsValidCode() error {
	if r.Codigo == "" {
		return errors.New("Por favor, ingresa un código válido")
	}
	if strings.Contains(r.Codigo, " ") {
		return errors.New("Por favor, asegúrate de que el código no contenga espacios en blanco")
	}
	matched, err := regexp.MatchString("^[a-zA-Z0-9]{6}$", r.Codigo)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("Por favor, asegúrate de que el código contenga solo letras y números")
	}
	if len(r.Codigo) != 6 {
		return errors.New("Por favor, asegúrate de que el código tenga exactamente 6 caracteres")
	}
	return nil
}
