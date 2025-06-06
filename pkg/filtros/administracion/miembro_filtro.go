package filtros

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros"

type MiembroFiltro struct {
	filtros.Paginacion
	Nombre                 string
	Apellido               string
	ComunidadID            uint
	SexoMasculino          bool
	SexoFemenino           bool
	OrdenarFechaNacimiento bool
	UsuarioID              uint
	Activos                bool
	Expulsados             bool
	SoloMiembros           bool
	SoloAdministradores    bool
	AdministradorID        uint
}
