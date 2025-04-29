package entities

import (
	"gorm.io/gorm"
)

type Comunidad struct {
	gorm.Model
	Nombre          string        `json:"nombre"`
	Descripcion     string        `json:"descripcion"`
	CodigoAcceso    string        `json:"codigo_acceso"`
	Habilitada      bool          `json:"habilitada"`
	FotoPerfil      string        `json:"foto_perfil"`
	LocalidadesId   int64         `json:"localidades_id"`
	TipoComunidadId int64         `json:"tipo_comunidad_id"`
	Email           string        `json:"email"`
	Telefono        string        `json:"telefono"`
	Cuit            string        `json:"cuit"`
	WebUrl          string        `json:"web_url"`
	Calle           string        `json:"calle"`
	Altura          int           `json:"altura"`
	NumeroPiso      uint          `json:"numero_piso"`
	Lat             float64       `json:"lat"`
	Lng             float64       `json:"lng"`
	Usuarios        []Usuario     `gorm:"many2many:usuarios_has_comunidades;"` // Relación de muchos a muchos
	Localidad       Localidad     `gorm:"foreignKey:LocalidadesId"`
	TipoComunidad   TipoComunidad `gorm:"foreignKey:TipoComunidadId"`
}

func (Comunidad) TableName() string {
	return "comunidades"
}
