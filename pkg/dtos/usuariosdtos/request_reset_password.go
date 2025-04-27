package usuariosdtos

import (
	"errors"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/commons"
)

type RequestResetPassword struct {
	NewPassword       string `json:"new_password"`
	RepeatNewPassword string `json:"repeat_new_password"`
	Token             string `json:"token"`
}

func (r *RequestResetPassword) Validate() error {
	if commons.StringIsEmpty(r.Token) {
		return errors.New("Datos incompletos")
	}

	err := commons.ValidatePassword(r.NewPassword)
	if err != nil {
		return err
	}
	err = commons.ValidatePassword(r.RepeatNewPassword)
	if err != nil {
		return err
	}

	return nil
}
