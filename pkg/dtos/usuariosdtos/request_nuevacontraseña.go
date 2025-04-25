package usuariosdtos

import (
	"errors"
	"strings"
)

type RequestChangePassword struct {
	Password          string `json:"password"`
	NewPassword       string `json:"new_password"`
	RepeatNewPassword string `json:"repeat_new_password"`
}

func (req *RequestChangePassword) Validar() error {

	// Verificar que las contraseñas no estén vacías
	if req.Password == "" || req.RepeatNewPassword == "" || req.NewPassword == "" {
		return errors.New("Las contraseñas no pueden estar vacías")
	}

	// Verifica que la nueva contraseña no contenga espacios en blanco
	if strings.Contains(req.NewPassword, " ") {
		return errors.New("La nueva contraseña no puede contener espacios en blanco")
	}

	if req.NewPassword != req.RepeatNewPassword {
		return errors.New("Las contraseñas no coinciden")
	}

	if req.Password == req.NewPassword {
		return errors.New("La nueva contraseña no puede ser igual a la contraseña actual")
	}

	if len(req.NewPassword) < 8 {
		return errors.New("La nueva contraseña debe tener al menos 8 caracteres")
	}

	return nil
}

func (req *RequestChangePassword) ValidarResetPassword() error {

	// Verificar que las contraseñas no estén vacías
	if req.RepeatNewPassword == "" || req.NewPassword == "" {
		return errors.New("Las contraseñas no pueden estar vacías")
	}

	// Verifica que la nueva contraseña no contenga espacios en blanco
	if strings.Contains(req.NewPassword, " ") {
		return errors.New("La nueva contraseña no puede contener espacios en blanco")
	}

	if req.NewPassword != req.RepeatNewPassword {
		return errors.New("Las contraseñas no coinciden")
	}

	if req.Password == req.NewPassword {
		return errors.New("La nueva contraseña no puede ser igual a la contraseña actual")
	}

	if len(req.NewPassword) < 8 {
		return errors.New("La nueva contraseña debe tener al menos 8 caracteres")
	}

	return nil
}
