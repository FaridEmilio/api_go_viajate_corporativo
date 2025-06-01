package filtros

type TrayectoFiltro struct {
	Paginacion
	ID              uint
	UsuariosID      uint
	ComunidadID     uint
	OriginName      string
	DestinationName string
	FechaInicio     string
	FechaFin        string
}
