package filtros

import "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros"

type UsuarioFiltro struct {
	filtros.Paginacion
	Email                 string
	Telefono              string
	Nombre                string
	Apellido              string
	RolesID               int64
	ID                    uint
	IDs                   []uint
	CargarPermisos        bool
	CargarComunidades     bool
	ComunidadSelectFields []string
	SelectFields          []string
	ComunidadID           uint
}
