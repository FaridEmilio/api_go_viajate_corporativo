package entities

import (
	"fmt"
	"strings"

	"github.com/faridEmilio/api_go_gym_manager/pkg/dtos/tools"
	"gorm.io/gorm"
)

type Configuracione struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Valor       string `json:"valor"`
}

func (c *Configuracione) IsValid() error {

	if tools.EsStringVacio(c.Nombre) {
		return fmt.Errorf("el campo nombre es obligatorio")
	}
	if tools.EsStringVacio(c.Valor) {
		return fmt.Errorf("el campo valor es obligatorio")
	}

	c.Nombre = strings.ToUpper(c.Nombre)
	c.Nombre = strings.TrimSpace(c.Nombre)

	return nil
}
