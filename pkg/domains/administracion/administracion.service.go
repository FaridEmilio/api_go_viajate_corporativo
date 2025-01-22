package administracion

import (
	"math"

	"github.com/faridEmilio/api_go_gym_manager/pkg/commons"
	"github.com/faridEmilio/api_go_gym_manager/pkg/domains/util"
	"github.com/faridEmilio/api_go_gym_manager/pkg/dtos"
)

type service struct {
	repository Repository
	util       util.UtilService
	commons    commons.Commons
}

func NewAdministracionService(repo Repository, util util.UtilService, commons commons.Commons, firebaseRemoteRepo storage.FirebaseRemoteRepository) AdministracionService {
	return &service{
		repository: repo,
		util:       util,
		commons:    commons,
	}
}

// PAGINACION
func _setPaginacion(number uint32, size uint32, total int64) (meta dtos.Meta) {
	from := (number - 1) * size
	lastPage := math.Ceil(float64(total) / float64(size))

	meta = dtos.Meta{
		Page: dtos.Page{
			CurrentPage: int32(number),
			From:        int32(from),
			LastPage:    int32(lastPage),
			PerPage:     int32(size),
			To:          int32(number * size),
			Total:       int32(total),
		},
	}
	return
}

// func (s *service) GetMiembrosService(filtro filtros.MiembroFiltro) (response administraciondtos.ResponseMiembros, erro error) {
// 	miembros, total, erro := s.repository.GetMiembrosRepository(filtro)
// 	if erro != nil {
// 		return
// 	}
// 	// Seteo paginación
// 	if filtro.Number > 0 && filtro.Size > 0 {
// 		response.Meta = _setPaginacion(filtro.Number, filtro.Size, total)
// 	}
// 	response.FromEntities(miembros)
// 	return
// }
