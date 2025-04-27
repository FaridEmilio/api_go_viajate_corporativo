package comunidad

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"time"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/storage"
)

type ComunidadService interface {
	// Permisos
	//GetMisComunidadesService(filtro filtros.ComunidadFiltro) (response comunidaddtos.ResponseComunidades, erro error)

	// CRUD TRAYECTO
	// PostTrayectoService(request comunidaddtos.RequestTrayecto) (erro error)
	// GetTrayectosService(filtro filtros.TrayectoFiltro) (response comunidaddtos.ResponseTrayectos, erro error)

	// // Alta de un miembro en una comunidad
	// PostAltaMiembroService(request comunidaddtos.RequestAltaMiembro) (nombreComunidad string, erro error)
	GetComunidadesService(filtro comunidaddtos.RequestComunidad) (response comunidaddtos.ResponseComunidades, erro error)
	PostComunidadService(request comunidaddtos.RequestComunidad) (erro error)
	UploadImageToFirebase(file *multipart.FileHeader) (string, error)
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

// func (s *comunidadService) PostTrayectoService(request comunidaddtos.RequestTrayecto) (erro error) {

// 	erro = request.Validate()
// 	if erro != nil {
// 		return
// 	}

// 	trayecto := request.ToEntity()

// 	// genero un identificador unico universal al trayecto
// 	uuid := s.commons.NewUUID()
// 	trayecto.Uuid = uuid

// 	erro = s.repository.PostTrayectoRepository(trayecto)
// 	if erro != nil {
// 		return
// 	}

// 	return
// }

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
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(b)
}
