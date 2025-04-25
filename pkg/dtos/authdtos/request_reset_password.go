package authdtos

import (
	"fmt"

)

type RequestResetPassword struct {
	Token             string `json:"token"`
	NewPassword       string `json:"new_password"`
	RepeatNewPassword string `json:"repeat_new_password"`
}

func (r *RequestResetPassword) Validate() error {
	if commons.StringIsEmpty(r.Token) || commons.StringIsEmpty(r.NewPassword) || commons.StringIsEmpty(r.RepeatNewPassword) {
		return fmt.Errorf("Ingrese su nueva contraseña")
	}

	if commons.ContainsSpaces(r.NewPassword) || commons.ContainsSpaces(r.RepeatNewPassword) {
		return fmt.Errorf("Las contraseñas no pueden contener espacios en blanco")
	}

	err := commons.ValidatePassword(r.NewPassword)
	if err != nil {
		return err
	}
	err = commons.ValidatePassword(r.RepeatNewPassword)
	if err != nil {
		return err
	}

	if r.NewPassword != r.RepeatNewPassword {
		return fmt.Errorf("Las contraseñas no coinciden")
	}
	return nil
}
