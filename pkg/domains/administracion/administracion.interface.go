package administracion

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/administraciondtos"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/administracion"
)

type AdministracionService interface {
	GetPaisesService(filtro filtros.PaisFiltro) (response administraciondtos.ResponsePaises, erro error)
	//GetMiembrosService(filtro filtros.MiembroFiltro) (response administraciondtos.ResponseMiembros, erro error)
}
