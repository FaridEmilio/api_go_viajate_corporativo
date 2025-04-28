package entities

type UsuariosHasComunidades struct {
	ComunidadesId uint `json:"comunidades_id"`
	UsuariosId    uint `json:"usuarios_id"`
	Activo        bool `json:"activo"`
}
