package filtros

type UsuarioFiltro struct {
	Email                 string
	Telefono              string
	ID                    uint
	IDs                   []uint
	CargarPermisos        bool
	CargarComunidades     bool
	ComunidadSelectFields []string
}
