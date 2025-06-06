package administracion

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/administraciondtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/authdtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/administracion"
	usuarioFiltros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/usuarios"
)

type AdministracionService interface {
	GetPaisesService(filtro filtros.PaisFiltro) (response administraciondtos.ResponsePaises, erro error)
	//GetMiembrosService(filtro filtros.MiembroFiltro) (response administraciondtos.ResponseMiembros, erro error)
	PutUsuarioHasComunidadService(request comunidaddtos.RequestAltaMiembro) (erro error)
	GetMiembrosService(filtro filtros.MiembroFiltro) (response administraciondtos.ResponseMiembros, erro error)
	GetSedesService(comunidadID uint) (response administraciondtos.ResponseSedes, erro error)
	CreateSedeService(request administraciondtos.RequestCreateSede) (erro error)
	UpdateSedeActivaService(request administraciondtos.RequestCreateSede) (erro error)

	GetUsuariosService(filtro usuarioFiltros.UsuarioFiltro) (response authdtos.ResponseUsuarios, erro error)
	GetUsuarioService(filtro usuarioFiltros.UsuarioFiltro) (response authdtos.ResponseUsuario, erro error)
}
