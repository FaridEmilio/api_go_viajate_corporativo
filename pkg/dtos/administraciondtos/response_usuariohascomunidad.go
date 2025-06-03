package administraciondtos

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
)

type ResponseComunidadMembers struct {
	Comunidad comunidaddtos.ResponseComunidad
	Usuarios  []comunidaddtos.ResponseUsuarioComunidad
}

func (r *ResponseComunidadMembers) FromEntities(entidades []entities.UsuariosHasComunidades) {
	if len(entidades) == 0 {
		return
	}
	var comunidadDTO comunidaddtos.ResponseComunidad
	comunidadDTO.FromEntity(entidades[0].Comunidad)
	r.Comunidad = comunidadDTO

	for _, e := range entidades {
		var usuarioDTO comunidaddtos.ResponseUsuarioComunidad
		usuarioDTO.FromEntity(e.Usuario)
		r.Usuarios = append(r.Usuarios, usuarioDTO)
	}
}
