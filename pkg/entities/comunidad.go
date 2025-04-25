package entities

import (
	"gorm.io/gorm"
)

type Comunidad struct {
	gorm.Model
	Nombre        string               `json:"nombre"`
	Descripcion   string               `json:"descripcion"`
	UsuariosRoles *UsuarioRolComunidad `json:"usuarios_roles" gorm:"foreignKey:comunidades_id"` // Relación con la tabla intermedia
	CodigoAcceso  string               `json:"codigo_acceso"`
	Habilitada    bool                 `json:"habilitada"`
	FotoPerfil    bool                 `json:"foto_perfil"`
}

func (Comunidad) TableName() string {
	return "comunidades"
}
