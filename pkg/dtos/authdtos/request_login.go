package authdtos

import (
	"errors"

)

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RequestLogin) Validate() error {
	if commons.StringIsEmpty(r.Email) {
		return errors.New("El email es obligatorio")
	}
	if !commons.IsEmailValid(r.Email) {
		return errors.New("El correo electrónico ingresado no es válido. Por favor, verifica e intenta nuevamente")
	}
	if commons.StringIsEmpty(r.Password) {
		return errors.New("La contraseña es obligatoria")
	}
	return nil
}
