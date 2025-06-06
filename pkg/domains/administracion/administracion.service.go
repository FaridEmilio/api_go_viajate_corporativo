package administracion

import (
	"errors"
	"math"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/administraciondtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/authdtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/administracion"
	usuariofiltros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/usuarios"
)

type service struct {
	repository Repository
	util       util.UtilService
}

func NewAdministracionService(repo Repository, util util.UtilService) AdministracionService {
	return &service{
		repository: repo,
		util:       util,
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

func (s *service) GetPaisesService(filtro filtros.PaisFiltro) (response administraciondtos.ResponsePaises, erro error) {
	entities, erro := s.repository.GetPaisesRepository(filtro)
	if erro != nil {
		return
	}

	response.FromEntities(entities)
	return
}

func (s *service) PutUsuarioHasComunidadService(request comunidaddtos.RequestAltaMiembro) (erro error) {
	erro = s.repository.UpdateUsuarioHasComunidadRepository(request)
	if erro != nil {
		return
	}
	return
}

func (s *service) GetMiembrosService(filtro filtros.MiembroFiltro) (response administraciondtos.ResponseMiembros, erro error) {
	entities, erro := s.repository.GetMiembrosRepository(filtro)
	if erro != nil {
		return
	}

	response.ToMiembrosResponse(entities)
	return
}

func (s *service) GetSedesService(comunidadID uint) (response administraciondtos.ResponseSedes, erro error) {
	entities, erro := s.repository.GetSedesRepository(comunidadID)
	if erro != nil {
		return
	}
	response.FromEntities(entities)
	return
}

func (s *service) CreateSedeService(request administraciondtos.RequestCreateSede) (erro error) {
	if err := request.Validate(); err != nil {
		return err
	}
	return s.repository.CreateSedeRepository(request)
}

func (s *service) UpdateSedeActivaService(request administraciondtos.RequestCreateSede) (erro error) {
	if request.Id == 0 {
		return errors.New("el id de la sede es obligatorio")
	}
	return s.repository.UpdateSedeActivaRepository(request.Id, request.Active)
}

func (s *service) GetUsuariosService(filtro usuariofiltros.UsuarioFiltro) (response authdtos.ResponseUsuarios, erro error) {
	usuarios, total, erro := s.repository.GetUsuariosRepository(filtro)
	if erro != nil {
		return
	}

	if filtro.Size > 0 && filtro.Number > 0 {
		response.Meta = _setPaginacion(uint32(filtro.Number), uint32(filtro.Size), total)
	}

	response.ToResponseUsuarios(usuarios)
	return
}

func (s *service) GetUsuarioService(filtro usuariofiltros.UsuarioFiltro) (response authdtos.ResponseUsuario, erro error) {
	user, erro := s.repository.GetUsuarioRepository(filtro)
	if erro != nil {
		return
	}

	response.FromEntity(user)
	return
}
