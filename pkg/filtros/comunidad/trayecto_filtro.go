package filtros

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros"

type TrayectoFiltro struct {
	filtros.Paginacion
	ID              uint
	UsuariosID      uint
	ComunidadID     uint
	OriginName      string
	DestinationName string
	FechaInicio     string
	FechaFin        string
}
