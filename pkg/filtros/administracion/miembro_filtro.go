package filtros

type MiembroFiltro struct {
	Paginacion
	Nombre                    string
	Apellido                  string
	ComunidadID               uint
	SexoMasculino             bool
	SexoFemenino              bool
	OrdenarFechaNacimiento    bool
	CargarSoloMiembros        bool
	CargarSoloAdministradores bool
	UsuarioID                 uint
	CargarEliminados          bool
}
