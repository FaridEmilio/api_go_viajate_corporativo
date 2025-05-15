package entities

type UsuariosHasComunidades struct {
	UsuariosID    uint `gorm:"column:usuarios_id;primaryKey"`
	ComunidadesID uint `gorm:"column:comunidades_id;primaryKey"`
	Activo        bool `gorm:"column:activo;default:true;not null"`
}

// Esto asegura que GORM use el nombre correcto para la tabla
func (UsuariosHasComunidades) TableName() string {
	return "usuarios_has_comunidades"
}
