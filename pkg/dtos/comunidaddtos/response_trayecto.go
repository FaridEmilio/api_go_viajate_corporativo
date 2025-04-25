package comunidaddtos

import (
	"github.com/faridEmilio/api_go_viajate/pkg/dtos"
	"github.com/faridEmilio/api_go_viajate/pkg/entities"
)

type ResponseTrayectos struct {
	Trayectos []ResponseTrayecto `json:"trayectos"`
	Meta      dtos.Meta          `json:"meta"`
	Total     int                `json:"total"`
}

type ResponseTrayecto struct {
	ID          uint                     `json:"id"`
	Uuid        string                   `json:"uuid"`
	Descripcion string                   `json:"descripcion"`
	Precio      uint                     `json:"precio"`
	Usuario     ResponseUsuarioComunidad `json:"usuario"`
	EstadosID   uint                     `json:"estados_id"`
	Estado      string                   `json:"estado"`
	Detalles    ResponseTrayectoDetalle  `json:"detalles"`
	Rutinas     []ResponseRutina         `json:"rutinas"`
}

type ResponseTrayectoDetalle struct {
	FrecuenciaSemanal bool `json:"frecuencia_semanal"`
	SoloEstudiantes   bool `json:"solo_estudiantes"`
	SoloMujeres       bool `json:"solo_mujeres"`
	Mascotas          bool `json:"mascotas"`
	Asientos          uint `json:"asientos"`
}

type ResponseRutina struct {
	ID    uint   `json:"id"`
	Dia   string `json:"dia"`
	Hora  string `json:"hora"`
	Fecha string `json:"fecha"`
}

func (r *ResponseTrayectos) ToResponseTrayectos(trayectos []*entities.Trayecto) {
	r.Total = len(trayectos)
	for _, trayecto := range trayectos {
		var temp ResponseTrayecto
		temp.ToResponseTrayecto(trayecto)
		r.Trayectos = append(r.Trayectos, temp)
	}
}

func (r *ResponseTrayecto) ToResponseTrayecto(t *entities.Trayecto) {
	r.ID = t.ID
	r.Uuid = t.Uuid
	r.Descripcion = t.Descripcion
	r.Precio = t.Precio
	r.EstadosID = t.EstadosID
	r.Usuario.FromEntity(t.Usuario)
	r.Estado = t.Estado.Nombre
	r.Detalles.ToResponseTrayectoDetalle(t.Detalles)
	r.ToResponseRutinas(t.Rutinas)
}

func (r *ResponseTrayectoDetalle) ToResponseTrayectoDetalle(entity *entities.TrayectoDetalle) {
	r.FrecuenciaSemanal = entity.FrecuenciaSemanal
	r.SoloEstudiantes = entity.SoloEstudiantes
	r.SoloMujeres = entity.SoloMujeres
	r.Mascotas = entity.Mascotas
	r.Asientos = entity.Asientos
}

func (r *ResponseRutina) ToResponseRutina(entity *entities.Rutina) {
	r.ID = entity.ID
	r.Dia = entity.Dia
	r.Hora = entity.Hora
	r.Fecha = entity.Fecha.Format("02-01-2006") // TODO formatear la salida de la fecha
}

func (r *ResponseTrayecto) ToResponseRutinas(rutinas []*entities.Rutina) {
	for _, rutina := range rutinas {
		var temp ResponseRutina
		temp.ToResponseRutina(rutina)
		r.Rutinas = append(r.Rutinas, temp)
	}
}
