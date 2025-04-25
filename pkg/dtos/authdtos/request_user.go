package authdtos

import (
	"errors"

)

type RequestUser struct {
	Name               string `json:"name"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	RepeatPassword     string `json:"repeat_password"`
	Telefono           string `json:"telefono"`
	Dni                string `json:"dni"`
	ProvincesID        int64  `json:"provinces_id"`
	RolesID            int64  `json:"roles_id"`
	CuilCuit           string `json:"cuil_cuit"`
	TermsAndConditions bool   `json:"terms_and_conditions"`
	AcceptPromotions   bool   `json:"accept_promotions"`
}

func (r *RequestUser) Validate() error {

	if commons.StringIsEmpty(r.Name) {
		return errors.New("El nombre es obligatorio")
	}
	if err := commons.IsNameValid(r.Name); err != nil {
		return err
	}

	if commons.StringIsEmpty(r.Email) {
		return errors.New("El email es obligatorio")
	}
	if !commons.IsEmailValid(r.Email) {
		return errors.New("El correo electrónico ingresado no es válido. Por favor, verifica e intenta nuevamente")
	}

	if commons.StringIsEmpty(r.Password) {
		return errors.New("La contraseña es obligatoria")
	}
	if commons.StringIsEmpty(r.RepeatPassword) {
		return errors.New("Vuelve a ingresar la contraseña por favor")
	}
	err := commons.ValidatePassword(r.Password)
	if err != nil {
		return err
	}
	err = commons.ValidatePassword(r.RepeatPassword)
	if err != nil {
		return err
	}
	if r.Password != r.RepeatPassword {
		return errors.New("Las contraseñas no coinciden")
	}

	if commons.StringIsEmpty(r.Telefono) {
		return errors.New("El número de teléfono es obligatorio")
	}
	if err := commons.IsValidPhoneNumber(r.Telefono); err != nil {
		return err
	}

	if !commons.IsDniValid(r.Dni) {
		return errors.New("El DNI es inválido")
	}

	// TODO validacion de cuil/cuit
	return nil

}

func (r *RequestUser) ToEntity() entities.Users {
	return entities.Users{
		Name:               r.Name,
		Email:              r.Email,
		Password:           r.Password,
		Telefono:           r.Telefono,
		Dni:                r.Dni,
		ProvincesID:        uint(r.ProvincesID),
		RolesID:            3, // rol de usuario por defecto
		CuilCuit:           r.CuilCuit,
		TermsAndConditions: r.TermsAndConditions,
		AcceptPromotions:   r.AcceptPromotions,
	}
}
