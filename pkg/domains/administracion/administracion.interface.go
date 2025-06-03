package administracion

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/administraciondtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/administracion"
)

type AdministracionService interface {
	GetPaisesService(filtro filtros.PaisFiltro) (response administraciondtos.ResponsePaises, erro error)
	//GetMiembrosService(filtro filtros.MiembroFiltro) (response administraciondtos.ResponseMiembros, erro error)
	PutUsuarioHasComunidadService(request comunidaddtos.RequestAltaMiembro) (erro error)
	GetComunidadMembersService(comunidadID uint) (response administraciondtos.ResponseComunidadMembers, erro error)
	GetSedesService(comunidadID uint) (response administraciondtos.ResponseSedes, erro error)
	CreateSedeService(request administraciondtos.RequestCreateSede) (erro error)
	UpdateSedeActivaService(request administraciondtos.RequestCreateSede) (erro error)
}
