package comunidad

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"mime/multipart"
	"strings"
	"time"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/storage"
)

type ComunidadService interface {
	// Permisos
	//GetMisComunidadesService(filtro filtros.ComunidadFiltro) (response comunidaddtos.ResponseComunidades, erro error)

	// CRUD TRAYECTO
	PostTrayectoService(request comunidaddtos.RequestTrayecto) (erro error)
	// GetTrayectosService(filtro filtros.TrayectoFiltro) (response comunidaddtos.ResponseTrayectos, erro error)

	// // Alta de un miembro en una comunidad
	GetComunidadesService(filtro comunidaddtos.RequestComunidad) (response comunidaddtos.ResponseComunidades, erro error)
	PostComunidadService(request comunidaddtos.RequestComunidad) (erro error)
	UploadImageToFirebase(file *multipart.FileHeader) (string, error)
	PutComunidadService(request comunidaddtos.RequestComunidad) (erro error)
	PostUsuarioComunidadService(request comunidaddtos.RequestAltaMiembro) (nombreComunidad string, erro error)
	PutUsuarioComunidadService(request comunidaddtos.RequestAltaMiembro) (erro error)
	GetTipoComunidadService(request comunidaddtos.RequestTipoComunidad) (response comunidaddtos.ResponseTipoComunidades, erro error)
}

func NewComunidadService(repo ComunidadRepository, util util.UtilService, firebaseRepo storage.FirebaseRemoteRepository) ComunidadService {
	service := comunidadService{
		repository:               repo,
		util:                     util,
		firebaseRemoteRepository: firebaseRepo,
	}
	return &service
}

type comunidadService struct {
	repository               ComunidadRepository
	util                     util.UtilService
	firebaseRemoteRepository storage.FirebaseRemoteRepository
}

// func (s *comunidadService) PostAltaMiembroService(request comunidaddtos.RequestAltaMiembro) (nombreComunidad string, erro error) {

// 	erro = request.IsValidCode()
// 	if erro != nil {
// 		return
// 	}

// 	// Buscar la comunidad por codigo
// 	filtro := filtros.ComunidadFiltro{
// 		CodigoAcceso: request.Codigo,
// 	}

// 	comunidades, erro := s.repository.GetComunidadesRepository(filtro)
// 	if erro != nil || len(comunidades) < 1 {
// 		erro = fmt.Errorf("El código ingresado no está asociado a ninguna comunidad. Asegúrate de que sea el correcto")
// 		return
// 	}

// 	// Seteo la comunidad encontrada
// 	comunidad := comunidades[len(comunidades)-1]

// 	// Verificar INEXISTENCIA del usuario en la comunidad
// 	filtroComunidad := filtros.ComunidadFiltro{
// 		UsuarioID: request.UsuariosID,
// 		ID:        comunidad.ID,
// 	}

// 	isMember, _ := s.GetPermisoComunidadService(filtroComunidad)
// 	if isMember {
// 		erro = fmt.Errorf("¡Ya formas parte de esta comunidad!")
// 		return
// 	}

// 	// crear la relacion usuario comunidad rol
// 	usuarioRolComunidad := entities.UsuarioRolComunidad{
// 		UsuariosID:    int64(request.UsuariosID),
// 		ComunidadesID: comunidad.ID,
// 		Rol:           "MIEMBRO",
// 	}

// 	err := s.repository.PostUsuarioRolComunidadReporitory(usuarioRolComunidad)
// 	if err != nil {
// 		erro = fmt.Errorf("Hubo un error al registrarte en esta comunidad")
// 		return
// 	}

// 	nombreComunidad = comunidad.Nombre
// 	return
// }

// func (s *comunidadService) GetPermisoComunidadService(filtro filtros.ComunidadFiltro) (acceso bool, erro error) {
// 	acceso, erro = s.repository.GetPermisoComunidadesRepository(filtro)
// 	if erro != nil {
// 		return
// 	}
// 	return
// }

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

// func (s *comunidadService) GetMisComunidadesService(filtro filtros.ComunidadFiltro) (response comunidaddtos.ResponseComunidades, erro error) {

// 	comunidades, erro := s.repository.GetComunidadesRepository(filtro)
// 	if erro != nil {
// 		return
// 	}

// 	response.FromEntities(comunidades)
// 	return
// }

func (s *comunidadService) PostTrayectoService(request comunidaddtos.RequestTrayecto) (erro error) {

	// erro = request.Validate()
	// if erro != nil {
	// 	return
	// }

	trayecto := request.ToEntity()
	erro = s.repository.PostTrayectoRepository(trayecto)
	if erro != nil {
		return
	}

	return
}

// // CRUD TRAYECTO
// func (s *comunidadService) GetTrayectosService(filtro filtros.TrayectoFiltro) (response comunidaddtos.ResponseTrayectos, erro error) {

// 	filtro.CargarDetalle = true
// 	trayectos, total, erro := s.repository.GetTrayectosRepository(filtro)
// 	if erro != nil {
// 		return
// 	}
// 	if filtro.Number > 0 && filtro.Size > 0 {
// 		response.Meta = _setPaginacion(filtro.Number, filtro.Size, total)
// 	}
// 	response.ToResponseTrayectos(trayectos)
// 	return

// }

func (s *comunidadService) GetComunidadesService(request comunidaddtos.RequestComunidad) (response comunidaddtos.ResponseComunidades, erro error) {
	comunidades, total, erro := s.repository.GetComunidadesRepository(request)
	if erro != nil {
		return
	}

	response.FromEntities(comunidades)

	if request.Size > 0 && request.Number > 0 {
		response.Meta = _setPaginacion(uint32(request.Number), uint32(request.Size), total)
	}
	return
}

func (s *comunidadService) PostComunidadService(request comunidaddtos.RequestComunidad) (erro error) {

	erro = request.Validate()
	if erro != nil {
		return
	}
	requesttipo := comunidaddtos.RequestTipoComunidad{
		Id: int(request.TipoComunidadId),
	}
	tipocomunidad, _, erro := s.repository.GetTipoComunidadRepository(requesttipo)
	if erro != nil {
		return
	}
	if len(tipocomunidad) < 1 {
		erro = errors.New("tipo de comunidad inexistente")
		return
	}
	comunidades, _, erro := s.repository.GetComunidadesRepository(request)
	if erro != nil {
		return
	}
	if len(comunidades) > 0 {
		erro = errors.New("comunidad existente")
		return
	}
	comunidad := request.ToEntity()
	uuid := NewUUID()
	comunidad.CodigoAcceso = uuid
	erro = s.repository.PostComunidadRepository(*comunidad)
	if erro != nil {
		return
	}
	return
}

func (s *comunidadService) UploadImageToFirebase(file *multipart.FileHeader) (string, error) {
	ctx := context.Background()

	openedFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer openedFile.Close()

	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	url, err := s.firebaseRemoteRepository.UploadFile(
		ctx,
		"comunidades",
		fileName,
		fileBytes,
		file.Header.Get("Content-Type"),
	)
	if err != nil {
		return "", err
	}

	return url, nil
}

func NewUUID() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	var result strings.Builder

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}
		result.WriteByte(charset[n.Int64()])
	}

	return result.String()
}

func (s *comunidadService) PutComunidadService(request comunidaddtos.RequestComunidad) (erro error) {
	if request.ID <= 0 {
		erro = errors.New("debe proporcionar un id para buscar la comunidad")
		return
	}
	comunidades, _, erro := s.repository.GetComunidadesRepository(request)
	if erro != nil {
		return
	}
	if len(comunidades) < 0 {
		erro = errors.New("comunidad inexistente")
		return
	}
	comunidad := comunidades[0]
	if len(request.Nombre) > 0 {
		comunidad.Nombre = request.Nombre
	}
	if len(request.Descripcion) > 0 {
		comunidad.Descripcion = request.Descripcion
	}
	if request.Habilitada != nil {
		comunidad.Habilitada = *request.Habilitada
	}

	erro = s.repository.UpdateComunidadRepository(comunidad)
	return
}

func (s *comunidadService) PostUsuarioComunidadService(request comunidaddtos.RequestAltaMiembro) (nombreComunidad string, erro error) {
	if len(request.Codigo) < 1 {
		erro = errors.New("debe proporcionar un codigo de comunidad")
		return
	}
	if request.UsuariosID < 1 {
		erro = errors.New("debe proporcionar un usuario")
		return
	}
	req := comunidaddtos.RequestComunidad{
		CodigoAcceso: request.Codigo,
	}
	comunidades, _, erro := s.repository.GetComunidadesRepository(req)
	if erro != nil {
		return
	}
	if len(comunidades) < 1 {
		erro = errors.New("comunidad inexistente")
		return
	}
	if comunidades[0].Habilitada == false {
		erro = errors.New("comunidad desactivada")
		return
	}

	entity := entities.UsuariosHasComunidades{
		ComunidadesID: comunidades[0].ID,
		UsuariosID:    request.UsuariosID,
	}

	erro = s.repository.PostUsuarioComunidadRepository(entity)
	if erro != nil {
		return
	}

	nombreComunidad = comunidades[0].Nombre
	return
}

func (s *comunidadService) PutUsuarioComunidadService(request comunidaddtos.RequestAltaMiembro) (erro error) {
	if request.Activo == nil {
		erro = errors.New("debe enviar una funcion")
		return
	}
	usuariocomunidad, erro := s.repository.GetUsuarioComunidadRepository(request)
	if erro != nil {
		return
	}
	if len(usuariocomunidad) < 1 {
		erro = errors.New("comunidad inexistente")
		return
	}
	if request.Activo != nil {
		usuariocomunidad[0].Activo = *request.Activo
	}

	// Ahora actualizamos en base de datos
	erro = s.repository.UpdateUsuarioComunidadRepository(usuariocomunidad[0])
	return
}

func (s *comunidadService) GetTipoComunidadService(request comunidaddtos.RequestTipoComunidad) (response comunidaddtos.ResponseTipoComunidades, erro error) {
	tipocomunidades, total, erro := s.repository.GetTipoComunidadRepository(request)
	if erro != nil {
		return
	}

	response.FromEntitiesTipoComunidad(tipocomunidades)

	if request.Size > 0 && request.Number > 0 {
		response.Meta = _setPaginacion(uint32(request.Number), uint32(request.Size), total)
	}
	return
}
