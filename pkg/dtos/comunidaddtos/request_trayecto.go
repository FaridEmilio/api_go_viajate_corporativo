package comunidaddtos

import (
	"errors"
	"time"

	"github.com/faridEmilio/api_go_viajate/pkg/commons"
	"github.com/faridEmilio/api_go_viajate/pkg/dtos/rutadtos"
	"github.com/faridEmilio/api_go_viajate/pkg/entities"
)

type RequestTrayecto struct {
	Descripcion       string               `json:"descripcion"`
	Precio            uint                 `json:"precio"`
	UsuariosID        uint                 `json:"usuarios_id"`
	ComunidadesID     uint                 `json:"comunidades_id"`
	FrecuenciaSemanal bool                 `json:"frecuencia_semanal"`
	SoloEstudiantes   bool                 `json:"solo_estudiantes"`
	SoloMujeres       bool                 `json:"solo_mujeres"`
	Mascotas          bool                 `json:"mascotas"`
	Asientos          uint                 `json:"asientos"`
	Rutinas           []RequestRutina      `json:"rutinas"`
	Ruta              rutadtos.RequestRuta `json:"ruta"`
}

type RequestRutina struct {
	Dia   EnumDia   `json:"día"`
	Hora  string    `json:"hora"`
	Fecha time.Time `json:"fecha"`
}

// NOTE: se envia el valor "trayectoSemanal" para validar la fecha de los dias de la rutina,
// Si el trayecto es semanal no hace falta especificar los días
func (r *RequestRutina) Validate(TrayectoSemanal bool) error {
	if commons.StringIsEmpty(string(r.Dia)) {
		return errors.New("El campo día es obligatorio")
	}
	if err := r.Dia.IsDayValid(); err != nil {
		return err
	}
	if commons.StringIsEmpty(r.Hora) {
		return errors.New("El campo hora es obligatorio")
	}
	if !TrayectoSemanal {
		if r.Fecha.IsZero() {
			return errors.New("El campo fecha es obligatorio")
		}
	}
	return nil
}

type EnumDia string

// Constantes para los días de la semana
const (
	Lunes     EnumDia = "Lunes"
	Martes    EnumDia = "Martes"
	Miercoles EnumDia = "Miércoles"
	Jueves    EnumDia = "Jueves"
	Viernes   EnumDia = "Viernes"
	Sabado    EnumDia = "Sábado"
	Domingo   EnumDia = "Domingo"
)

// Función para validar si el día es válido
func (enumDia EnumDia) IsDayValid() error {
	switch enumDia {
	case Lunes, Martes, Miercoles, Jueves, Viernes, Sabado, Domingo:
		return nil
	}
	return errors.New("Día con formato inválido")
}

func (r *RequestTrayecto) Validate() error {
	if r.UsuariosID <= 0 {
		return errors.New("Usuario no encontrado")
	}
	if r.Precio <= 0 {
		return errors.New("El precio debe ser mayor que cero")
	}
	// Validaciones de los detalles
	if r.Asientos <= 0 || r.Asientos > 6 {
		return errors.New("El número de asientos es inválido")
	}
	// Validaciones para cada rutina/frecuencia del request
	for _, rutina := range r.Rutinas {
		if err := rutina.Validate(r.FrecuenciaSemanal); err != nil {
			return err
		}
	}

	// Validar la Ruta
	if err := r.Ruta.Validate(); err != nil {
		return err
	}

	return nil
}

func (r *RequestTrayecto) ToEntity() (trayecto entities.Trayecto) {
	trayecto.UsuariosID = r.UsuariosID
	trayecto.ComunidadesID = r.ComunidadesID
	trayecto.Descripcion = r.Descripcion
	trayecto.Precio = r.Precio
	trayecto.EstadosID = 1 // 1 - Disponible
	trayecto.Detalles = r.ToTrayectoDetallesEntity()
	trayecto.Rutinas = r.ToEntityRutinas()
	trayecto.Ruta = r.Ruta.ToEntity()
	return
}

func (r *RequestTrayecto) ToTrayectoDetallesEntity() *entities.TrayectoDetalle {
	return &entities.TrayectoDetalle{
		FrecuenciaSemanal: r.FrecuenciaSemanal,
		SoloEstudiantes:   r.SoloEstudiantes,
		SoloMujeres:       r.SoloMujeres,
		Mascotas:          r.Mascotas,
		Asientos:          r.Asientos,
	}
}

func (r *RequestRutina) ToEntityRutina() *entities.Rutina {
	return &entities.Rutina{
		Dia:   string(r.Dia), // Formateo a string EnumDia
		Hora:  r.Hora,
		Fecha: r.Fecha,
	}
}

func (r *RequestTrayecto) ToEntityRutinas() []*entities.Rutina {
	rutinasEntities := make([]*entities.Rutina, len(r.Rutinas))
	for i, rutinaRequest := range r.Rutinas {
		rutinasEntities[i] = rutinaRequest.ToEntityRutina()
	}

	return rutinasEntities
}
