package authdtos

import (
	"errors"
	"strings"
	"unicode"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/commons"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type RequestNewUser struct {
	Nombre            string     `json:"nombre"`
	Apellido          string     `json:"apellido"`
	Email             string     `gorm:"unique;not null" json:"email"`
	Contraseña        string     `json:"contraseña"`
	Telefono          string     `json:"telefono"`
	FechaNacimiento   string     `json:"fecha_nacimiento"`
	Genero            EnumGenero `json:"genero"`
	RepetirContraseña string     `json:"repetir_contraseña"`
	Terminos          bool       `json:"terminos"`
}

type EnumGenero string

const (
	GeneroMasculino EnumGenero = "MASCULINO"
	GeneroFemenino  EnumGenero = "FEMENINO"
	GeneroOtro      EnumGenero = "OTRO"
)

// Validar valida todos los campos de RequestNewUser
func (req *RequestNewUser) Validate() error {

	// Verificar que las contraseñas no estén vacías
	if req.Contraseña == "" || req.RepetirContraseña == "" {
		return errors.New("Las contraseñas no pueden estar vacías")
	}

	if strings.TrimSpace(req.Nombre) == "" {
		return errors.New("El nombre es obligatorio")
	}
	if err := isValidName(req.Nombre); err != nil {
		return err
	}
	if strings.TrimSpace(req.Apellido) == "" {
		return errors.New("El apellido es obligatorio")
	}
	if err := isValidName(req.Apellido); err != nil {
		return err
	}
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("El email es obligatorio")
	}
	if err := isValidEmail(req.Email); err != nil {
		return err
	}
	if strings.TrimSpace(req.Contraseña) == "" {
		return errors.New("La contraseña es obligatoria")
	}
	if strings.TrimSpace(req.FechaNacimiento) == "" {
		return errors.New("La fecha de nacimiento es obligatoria")
	}
	if err := req.Genero.IsGeneroValid(); err != nil {
		return err
	}
	if req.Contraseña != req.RepetirContraseña {
		return errors.New("Las contraseñas no coinciden")
	}
	if strings.Contains(req.Contraseña, " ") || strings.Contains(req.RepetirContraseña, " ") {
		return errors.New("La contraseña no puede contener espacios en blanco")
	}
	if strings.TrimSpace(req.Telefono) == "" {
		return errors.New("El número de teléfono es obligatorio")
	}
	if !req.Terminos {
		return errors.New("Debe aceptar los términos y condiciones")
	}
	if err := isValidPhoneNumber(req.Telefono); err != nil {
		return err
	}
	err := commons.ValidatePassword(req.Contraseña)
	if err != nil {
		return err
	}
	err = commons.ValidatePassword(req.RepetirContraseña)
	if err != nil {
		return err
	}

	return nil
}

func isValidEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("El email debe contener un '@'")
	}
	return nil
}

func (enumGenero EnumGenero) IsGeneroValid() error {
	switch enumGenero {
	case GeneroMasculino, GeneroFemenino, GeneroOtro:
		return nil
	}
	return errors.New("Genero con formato inválido")
}

func isValidPhoneNumber(phoneNumber string) error {
	for _, char := range phoneNumber {
		if !unicode.IsDigit(char) {
			return errors.New("El número de teléfono solo puede contener caracteres numéricos")
		}
	}
	if len(phoneNumber) < 10 {
		return errors.New("El número de teléfono debe tener al menos 10 caracteres")
	}
	return nil
}
func FormatPhoneNumber(phoneNumber string) string {
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "+", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	if len(phoneNumber) > 0 && (phoneNumber[0] == '9' || phoneNumber[0] == '0') {
		phoneNumber = phoneNumber[1:]
	}
	if strings.HasPrefix(phoneNumber, "549") {
		phoneNumber = phoneNumber[3:]
	}
	if strings.HasPrefix(phoneNumber, "54") {
		phoneNumber = phoneNumber[2:]
	}
	return phoneNumber
}

// Formateo primer letra de un Nombre y Apellido
func FormatNombre(name string) string {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return name
	}
	return strings.ToUpper(string(name[0])) + name[1:]
}

// func createWaLink(phoneNumber string) string {
// 	formattedPhoneNumber := formatPhoneNumber(phoneNumber)
// 	return fmt.Sprintf("https://wa.me/549%s", formattedPhoneNumber)
// }

func (r *RequestNewUser) ToEntity() (UserEntity entities.Usuario) {
	UserEntity.Nombre = FormatNombre(r.Nombre)
	UserEntity.Apellido = FormatNombre(r.Apellido)
	UserEntity.Email = r.Email
	UserEntity.Telefono = FormatPhoneNumber(r.Telefono)
	UserEntity.Genero = string(r.Genero)
	UserEntity.Terminos = true
	UserEntity.RolesID = 3 // Por Defecto es Rol Usuario
	return
}
func isValidName(name string) error {
	if len(name) < 3 {
		return errors.New("El nombre y el apellido deben tener al menos 3 letras")
	}
	for _, char := range name {
		if !(unicode.IsLetter(char) || unicode.IsSpace(char)) {
			return errors.New("El nombre y el apellido solo pueden contener letras y espacios")
		}
	}
	return nil
}
