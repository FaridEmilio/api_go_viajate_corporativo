package entities

import (
	"gorm.io/gorm"
)

type Comunidad struct {
	gorm.Model
	Nombre          string    `json:"nombre"`
	Descripcion     string    `json:"descripcion"`
	Usuarios        []Usuario `gorm:"many2many:usuarios_has_comunidades;"` // Relación de muchos a muchos
	CodigoAcceso    string    `json:"codigo_acceso"`
	Habilitada      bool      `json:"habilitada"`
	FotoPerfil      string    `json:"foto_perfil"`
	LocalidadId     int64     `json:"localidad_id"`
	TipoComunidadId int64     `json:"tipo_comunidad_id"`
	Email           string    `json:"email"`
	Telefono        string    `json:"telefono"`
	Cuit            string    `json:"cuit"`
	WebUrl          string    `json:"web_url"`
	Calle           string    `json:"calle"`
	Altura          int       `json:"altura"`
	NumeroPiso      uint      `json:"numero_piso"`
	Lat             float32   `json:"lat"`
	Lng             float32   `json:"lng"`
}

func (Comunidad) TableName() string {
	return "comunidades"
}
