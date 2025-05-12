package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/commons"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/authdtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/usuarios"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterService(request authdtos.RequestNewUser) (response authdtos.ResponseLogin, erro error)
	LoginService(request authdtos.RequestLogin) (response authdtos.ResponseLogin, erro error)
	RefreshTokenService(userID uint) (response authdtos.ResponseLogin, erro error)
	GetTokensService(user entities.Usuario) (tokenResponse authdtos.ResponseLogin, erro error)

	// USUARIO
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

func (s *service) LoginService(request authdtos.RequestLogin) (response authdtos.ResponseLogin, erro error) {
	// Valido request
	erro = request.Validate()
	if erro != nil {
		return
	}

	// Verifico existencia de usuario registrado con email solicitado
	existsMail, erro := s.repository.GetUserExistsByEmail(request.Email)
	if erro != nil {
		return
	}

	if !existsMail {
		erro = fmt.Errorf("No existe una cuenta asociada con el correo electrónico proporcionado")
		return
	}

	// Recupero el usuario por email
	user, err := s.repository.FindByEmail(request.Email)
	if err != nil {
		return response, errors.New("No pudimos encontrar tu usuario. Por favor, revisa los datos e inténtalo de nuevo")
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Contraseña), []byte(request.Password)); err != nil {
		return response, errors.New("Contraseña incorrecta, inténtalo de nuevo")
	}

	response, err = s.GetTokensService(user)
	if err != nil {
		return
	}

	return
}

func (s *service) GetTokensService(user entities.Usuario) (tokenResponse authdtos.ResponseLogin, erro error) {
	rol := user.Rol.Rol
	permisos := make([]string, len(user.Rol.Permisos))
	for i, perm := range user.Rol.Permisos {
		permisos[i] = perm.Permiso
	}

	expiration := time.Now().Add(48 * time.Hour).Unix()
	userData := authdtos.ResponseUsuario{}
	userData.FromEntity(user)

	claims := jwt.MapClaims{
		"iss":      "Viajate",
		"sub":      fmt.Sprintf("%d", user.ID),
		"user":     userData,
		"exp":      expiration,
		"rol":      rol,
		"permisos": permisos,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return tokenResponse, fmt.Errorf("no se pudo firmar el token")
	}

	refreshTokenExpiration := time.Now().Add(15 * 24 * time.Hour)
	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "Viajate",
		Subject:   fmt.Sprintf("%d", user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Unix(refreshTokenExpiration.Unix(), 0)),
	})

	refreshToken, err := refreshTokenClaims.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return tokenResponse, fmt.Errorf("no se pudo generar el refresh token")
	}

	tokenResponse.Token = signedToken
	tokenResponse.RefreshToken = refreshToken
	return
}

func (s *service) RegisterService(request authdtos.RequestNewUser) (response authdtos.ResponseLogin, erro error) {
	emailValido := commons.IsEmailValid(request.Email)
	if !emailValido {
		erro = fmt.Errorf("El correo electrónico ingresado no es válido. Por favor, verifica e intenta nuevamente")
		return
	}

	// Verifico existencia de usuario registrado con email solicitado
	existsMail, erro := s.repository.GetUserExistsByEmail(request.Email)
	if erro != nil {
		return
	}

	if existsMail {
		erro = fmt.Errorf("El correo electrónico ingresado ya está registrado")
		return
	}

	erro = request.Validate()
	if erro != nil {
		erro = fmt.Errorf(erro.Error())
		return
	}

	fechaNacimiento, erro := commons.IsAdult(request.FechaNacimiento)
	if erro != nil {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Contraseña), 14)
	if err != nil {
		erro = fmt.Errorf("No se pudo completar el registro con éxito")
		return
	}

	user := request.ToEntity()
	user.Contraseña = string(hashedPassword)
	user.FechaNacimiento = fechaNacimiento
	userEntity, erro := s.repository.PostUsuarioRepository(user)
	if erro != nil {
		erro = fmt.Errorf("Error al crear usuario: " + erro.Error())
		return
	}

	if userEntity.ID < 1 {
		erro = fmt.Errorf("No se pudo completar el registro con éxito")
		return
	}

	// requestMailVerify := viajatedtos.RequestUserEmailVerification{
	// 	ID:    userEntity.ID,
	// 	Email: userEntity.Email,
	// 	Name:  userEntity.Nombre,
	// }

	// erro = s.GenerateAndSendEmailVerificationCode(requestMailVerify, runEndpoint)
	// if erro != nil {
	// 	return
	// }

	response, err = s.GetTokensService(userEntity)
	if err != nil {
		return
	}

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

func (s *service) GetUserService(filter filtros.UsuarioFiltro) (response authdtos.ResponseUsuario, erro error) {
	user, erro := s.repository.GetUserRepository(filter, nil)
	if erro != nil {
		return
	}

	response.FromEntity(user)
	return
}

func (s *service) RefreshTokenService(userID uint) (response authdtos.ResponseLogin, erro error) {
	if userID < 1 {
		erro = fmt.Errorf("Usuario no especificado")
		return
	}

	filtro := filtros.UsuarioFiltro{
		ID:             userID,
		CargarPermisos: true,
	}

	user, erro := s.repository.GetUserRepository(filtro, nil)
	if erro != nil {
		return
	}

	response, erro = s.GetTokensService(user)
	if erro != nil {
		return
	}

	return
}
