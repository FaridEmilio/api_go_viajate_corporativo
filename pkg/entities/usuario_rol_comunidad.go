package entities

import (
	"gorm.io/gorm"
)

type UsuarioRolComunidad struct {
	gorm.Model
	UsuariosID    int64      `json:"usuarios_id"`
	ComunidadesID uint       `json:"comunidades_id"`
	Usuario       *Usuario   `json:"usuario" gorm:"foreignKey:UsuariosID"`
	Comunidad     *Comunidad `json:"comunidad" gorm:"foreignKey:ComunidadesID"`
	Rol           EnumRoles  `json:"rol" gorm:"type:enum('ADMINISTRADOR', 'MIEMBRO');not null"`
}

func (UsuarioRolComunidad) TableName() string {
	return "usuarios_roles_comunidades"
}

type EnumRoles string

// ENUM DE ROLES DE PERMISOS
const (
	RolAdministrador EnumRoles = "ADMINISTRADOR"
	RolMiembro       EnumRoles = "MIEMBRO"
)
