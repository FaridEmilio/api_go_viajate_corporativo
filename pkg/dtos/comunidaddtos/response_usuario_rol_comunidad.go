package comunidaddtos

type ResponseUsuarioRolComunidad struct {
	ComunidadesID   uint   `json:"comunidades_id"`
	ComunidadNombre string `json:"comunidad_nombre"` // Nombre de la comunidad
	Rol             string `json:"rol"`              // rol (ADMINISTRADOR, MIEMBRO)
}
