package comunidaddtos

import (
	"errors"
	"fmt"
	"strings"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/commons"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type RequestVehiculo struct {
	MarcasID   uint   `json:"marcas_id"`
	UsuariosID uint   `json:"usuarios_id"`
	Modelo     string `json:"modelo"`
	Año        int64  `json:"año"`
	Color      string `json:"color"`
	Patente    string `json:"patente"`
	Tipo       string `json:"tipo"`
}

func (r *RequestVehiculo) ToEntity() entities.Vehiculo {
	return entities.Vehiculo{
		MarcasID:   r.MarcasID,
		UsuariosID: r.UsuariosID,
		Modelo:     r.Modelo,
		Año:        r.Año,
		Color:      r.Color,
		Patente:    r.Patente,
		Tipo:       entities.EnumTipoVehiculo(r.Tipo),
	}
}

func (r *RequestVehiculo) FormatVehiculoRequest() {
	r.Modelo = commons.FormatNombre(r.Modelo)
	r.Color = commons.FormatNombre(r.Color)
	r.Patente = strings.ToUpper(strings.Join(strings.Fields(r.Patente), ""))
}

func (r *RequestVehiculo) Validate() error {
	if r.MarcasID < 1 {
		return fmt.Errorf("Debes seleccionar la marca del vehículo")
	}
	if r.UsuariosID < 1 {
		return fmt.Errorf("usuario válido")
	}
	if commons.StringIsEmpty(r.Modelo) {
		return fmt.Errorf("Debes especificar el modelo del vehículo")
	}
	if commons.StringIsEmpty(r.Tipo) {
		return fmt.Errorf("Debes especificar el tipo de vehículo")
	}
	if r.Año <= 1900 || r.Año > 2100 {
		return fmt.Errorf("Debe ingresar un año válido para el vehículo")
	}
	if commons.StringIsEmpty(r.Color) {
		return fmt.Errorf("Debe especificar el color del vehículo")
	}
	if commons.StringIsEmpty(r.Patente) {
		return fmt.Errorf("La patente del vehículo no puede estar vacía")
	}
	if !commons.IsPatenteValida(r.Patente) {
		return errors.New("La patente es inválida")

	}

	return nil
}
