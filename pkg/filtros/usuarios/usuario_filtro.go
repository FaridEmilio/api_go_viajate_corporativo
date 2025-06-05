package filtros

type UsuarioFiltro struct {
	Email                 string
	Telefono              string
	Nombre                string
	Apellido              string
	ID                    uint
	IDs                   []uint
	CargarPermisos        bool
	CargarComunidades     bool
	ComunidadSelectFields []string
	ComunidadID           uint
	Activos               bool
	Expulsados            bool
	SoloMiembros          bool
	SoloAdministradores   bool
}
