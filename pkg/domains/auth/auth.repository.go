package auth

import (
	"errors"

	"github.com/faridEmilio/api_go_viajate_corporativo/internal/database"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/usuarios"
)

type Repository interface {
	GetUserRepository(filter filtros.UsuarioFiltro, fields []string) (user entities.Usuario, erro error)

	UpdateUserDataRepository(userID uint, updateFields map[string]interface{}) (erro error)
	GetUserByIDRepository(userID uint, fields []string) (user entities.Usuario, erro error)
	GetUsuariosRepository(filtro filtros.UsuarioFiltro) (usuarios []entities.Usuario, erro error)

	// Email verification
	CreateTokenEmailVerification(entity entities.EmailToken) error
	FindUserIDByEmailToken(token string) (userID uint, erro error)
	DeleteEmailTokenRepository(token string) error

	// RestorePassword
	CreatePasswordResetRepository(entity entities.PasswordReset) error
	FindUserIDByToken(token string) (userID uint, erro error)
	DeletePasswordResetByToken(token string) error
}

func NewAuthRepository(conn *database.MySQLClient, util util.UtilService) Repository {
	return &repository{
		SQLClient:   conn,
		utilService: util,
	}
}

type repository struct {
	SQLClient   *database.MySQLClient
	utilService util.UtilService
}

func (r *repository) UpdateUserDataRepository(id uint, updateFields map[string]interface{}) (erro error) {
	if err := r.SQLClient.Model(&entities.Usuario{}).
		Where("id = ?", id).
		Updates(updateFields).Error; err != nil {
		return err
	}
	return
}

func (r *repository) GetUserByIDRepository(userID uint, fields []string) (user entities.Usuario, erro error) {
	query := r.SQLClient.Model(entities.Usuario{})

	// Si se especifican campos, usar Select para elegirlos
	if len(fields) > 0 {
		query = query.Select(fields)
	}

	resp := query.Where("id = ?", userID).First(&user)
	if resp.Error != nil {
		erro = errors.New("usuario no encontrado")
		return
	}
	return
}

func (r *repository) GetUsuariosRepository(filtro filtros.UsuarioFiltro) (usuarios []entities.Usuario, erro error) {
	resp := r.SQLClient.Model(entities.Usuario{})

	if len(filtro.IDs) > 0 {
		resp.Where("usuarios_id IN (?)", filtro.IDs)
	}

	resp.Find(&usuarios)
	if resp.Error != nil {
		erro = errors.New("error al obtener reseñas")
	}
	return
}

// ************* Email verification *************
func (r *repository) CreateTokenEmailVerification(entity entities.EmailToken) (erro error) {
	resp := r.SQLClient.Create(&entity)
	if resp.Error != nil {
		return errors.New("error al crear email token")
	}
	return nil
}

func (r *repository) DeleteEmailTokenRepository(token string) (erro error) {
	return r.SQLClient.Where("token = ?", token).Delete(&entities.EmailToken{}).Error
}

func (r *repository) FindUserIDByEmailToken(token string) (userID uint, erro error) {
	err := r.SQLClient.Model(&entities.EmailToken{}).Where("token = ?", token).Pluck("usuarios_id", &userID).Error
	if err != nil {
		return
	}
	return
}

func (r *repository) GetUserRepository(filter filtros.UsuarioFiltro, fields []string) (user entities.Usuario, erro error) {
	resp := r.SQLClient.Model(entities.Usuario{})
	// Si se especifican campos, usar Select para elegirlos
	if len(fields) > 0 {
		resp = resp.Select(fields)
	}

	if len(filter.Email) > 0 {
		resp.Where("email = ?", filter.Email)
	}

	if filter.ID > 0 {
		resp.Where("id = ?", filter.ID)
	}

	resp.First(&user)
	if resp.Error != nil {
		erro = errors.New("user not found")
		return
	}
	return
}

func (r *repository) CreatePasswordResetRepository(entity entities.PasswordReset) error {
	resp := r.SQLClient.Create(&entity)
	if resp.Error != nil {
		return errors.New("error al crear password reset token")
	}
	return nil
}

func (r *repository) FindUserIDByToken(token string) (userID uint, erro error) {
	err := r.SQLClient.Model(&entities.PasswordReset{}).Where("token = ?", token).Pluck("usuarios_id", &userID).Error
	if err != nil {
		return
	}
	return
}

func (r *repository) DeletePasswordResetByToken(token string) error {
	return r.SQLClient.Where("token = ?", token).Delete(&entities.PasswordReset{}).Error
}
