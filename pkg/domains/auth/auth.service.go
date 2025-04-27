package auth

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/authdtos"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/usuarios"
)

type AuthService interface {
	Login(request authdtos.RequestLogin) (response authdtos.ResponseLogin, err error)
	GetUserService(filter filtros.UsuarioFiltro) (response authdtos.ResponseUsuario, erro error)
	//UpdateProfilePictureService(ctx context.Context, userID uint, fileHeader *multipart.FileHeader) (path string, err error)
	//DeleteProfilePictureService(ctx context.Context, userID uint) (erro error)
	//ChangePasswordService(userID uint, request usuariosdtos.RequestChangePassword) (erro error)
	//UpdateUserService(userID uint, request usuariosdtos.RequestUpdateUser) (erro error)
}

type service struct {
	repository Repository
	util       util.UtilService
}

func NewAuthService(repo Repository, util util.UtilService) AuthService {
	return &service{
		repository: repo,
		util:       util,
		//commons:                  commons,
		//firebaseRemoteRepository: firebaseRemoteRepo,
	}
}

// Login implements UsuarioService.
func (s *service) Login(request authdtos.RequestLogin) (response authdtos.ResponseLogin, err error) {
	return
}

// func (s *service) UpdateProfilePictureService(ctx context.Context, userID uint, fileHeader *multipart.FileHeader) (path string, err error) {

// 	// 1. Extraer la info del archivo subido y validar la imagen (para evitar lecturas innecesarias)
// 	fileData, fileExt, contentType, err := commons.IsProfilePhotoValid(fileHeader)
// 	if err != nil {
// 		return
// 	}

// 	// 2. Generar un UUID único para la imagen y nombre del archivo
// 	uuid := commons.NewUUID()
// 	filename := fmt.Sprintf("%s%s", uuid, fileExt)

// 	// 3. Subir la imagen a Firebase Storage
// 	path, err = s.firebaseRemoteRepository.UploadFile(ctx, "usuarios", filename, fileData, contentType)
// 	if err != nil {
// 		return
// 	}

// 	// 4. Actualizar la nueva URL al usuario
// 	updateData := map[string]interface{}{
// 		"foto_perfil": path,
// 	}

// 	err = s.repository.UpdateUserDataRepository(userID, updateData)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// func (s *service) DeleteProfilePictureService(ctx context.Context, userID uint) (erro error) {

// }

// func (s *service) ChangePasswordService(userID uint, request usuariosdtos.RequestChangePassword) (erro error) {

// 	// Se valida el nuevo request
// 	erro = request.Validar()
// 	if erro != nil {
// 		return
// 	}

// 	// Verificar que la contraseña ingresada sea igual a la del usuario
// 	fields := []string{"contraseña"} // solo select de contraseña
// 	user, erro := s.repository.GetUserByIDRepository(userID, fields)
// 	if erro != nil {
// 		return
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Contraseña), []byte(request.Password)); err != nil {
// 		erro = fmt.Errorf("La contraseña actual es incorrecta")
// 		return
// 	}

// 	// si todo es correcto, se hashea la nueva contraseña
// 	newPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 14)
// 	if err != nil {
// 		erro = fmt.Errorf("No se pudo actualizar la nueva contraseña")
// 		return
// 	}
// 	// 4. Actualizar entidad usuario
// 	updateData := map[string]interface{}{
// 		"contraseña": newPassword,
// 	}

// 	err = s.repository.UpdateUserDataRepository(userID, updateData)
// 	if err != nil {
// 		erro = fmt.Errorf("No se pudo actualizar la nueva contraseña")
// 		return
// 	}
// 	return
// }

// func (s *service) UpdateUserService(userID uint, request usuariosdtos.RequestUpdateUser) (erro error) {
// 	// Validaciones
// 	erro = request.Validate()
// 	if erro != nil {
// 		return
// 	}

// 	fechaNacimiento, erro := commons.IsAdult(request.FechaNacimiento)
// 	if erro != nil {
// 		return
// 	}

// 	updateData := map[string]interface{}{
// 		"nombre":           commons.FormatNombre(request.Nombre),
// 		"apellido":         commons.FormatNombre(request.Apellido),
// 		"fecha_nacimiento": fechaNacimiento,
// 		"genero":           string(request.Genero),
// 	}

// 	erro = s.repository.UpdateUserDataRepository(userID, updateData)
// 	if erro != nil {
// 		return
// 	}
// 	return
// }

// GetUserService implements AuthService.
func (s *service) GetUserService(filter filtros.UsuarioFiltro) (response authdtos.ResponseUsuario, erro error) {
	user, erro := s.repository.GetUserRepository(filter, nil)
	if erro != nil {
		return
	}

	response.FromEntity(user)
	return
}
