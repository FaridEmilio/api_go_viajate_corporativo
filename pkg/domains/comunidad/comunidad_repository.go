package comunidad

import (
	"errors"
	"fmt"

	"github.com/faridEmilio/api_go_viajate_corporativo/internal/database"
	"github.com/faridEmilio/api_go_viajate_corporativo/internal/logs"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/comunidad"
	"gorm.io/gorm"
)

type ComunidadRepository interface {
	GetComunidadesRepository(filtro filtros.ComunidadFiltro) (comunidades []entities.Comunidad, erro error)
	// CONSULTA PARA MIDDLEWARE
	GetPermisoComunidadesRepository(filtro filtros.ComunidadFiltro) (bool, error)

	// CRUD TRAYECTO
	PostTrayectoRepository(trayecto entities.Trayecto) error
	GetTrayectosRepository(filtro filtros.TrayectoFiltro) (trayectos []*entities.Trayecto, totalFilas int64, erro error)
	UpdateTrayectoRepository(id uint, updateFields map[string]interface{}) (erro error)

	//CRUD USUARIO_ROL_COMUNIDAD
	PostUsuarioRolComunidadReporitory(entity entities.UsuarioRolComunidad) (erro error)

	GetDB() (db *database.MySQLClient)
}

func NewComunidadRepository(conn *database.MySQLClient, util util.UtilService) ComunidadRepository {
	return &comunidadRepository{
		SqlClient:   conn,
		utilService: util,
	}
}

type comunidadRepository struct {
	SqlClient   *database.MySQLClient
	utilService util.UtilService
}

func (r *comunidadRepository) GetDB() (db *database.MySQLClient) {
	return r.SqlClient
}
func (r *comunidadRepository) GetPermisoComunidadesRepository(filtro filtros.ComunidadFiltro) (bool, error) {
	var exists bool
	resp := r.SqlClient.Model(&entities.Comunidad{})

	if (filtro.UsuarioID > 0) && (filtro.ID > 0) {
		resp = resp.Joins("INNER JOIN usuarios_roles_comunidades urc ON urc.comunidades_id = comunidades.id").
			Where("urc.usuarios_id = ? AND urc.comunidades_id = ? AND urc.deleted_at IS NULL", filtro.UsuarioID, filtro.ID)
	}

	if len(filtro.UsuarioScope) > 0 {
		resp.Where("urc.rol = ?", filtro.UsuarioScope)
	}

	resp.Select("1").Limit(1).Scan(&exists) // Usamos EXISTS para verificar si al menos un registro cumple la condición
	// Si exists es true, significa que existe al menos un registro que cumple la condición
	if exists {
		return true, nil // Middleware válido
	}

	return false, nil // Middleware no válido
}

func (r *comunidadRepository) GetComunidadesRepository(filtro filtros.ComunidadFiltro) (comunidades []entities.Comunidad, erro error) {
	resp := r.SqlClient.Model(&entities.Comunidad{})

	// FILTROS
	if filtro.ID > 0 {
		resp.Where("id = ?", filtro.ID).Limit(1)
	}
	if filtro.UsuarioID > 0 {
		resp = resp.Joins("LEFT JOIN usuarios_roles_comunidades urc ON urc.comunidades_id = comunidades.id").
			Where("urc.usuarios_id = ?", filtro.UsuarioID).Preload("UsuariosRoles", "usuarios_id = ?", filtro.UsuarioID)
	}
	if len(filtro.Nombre) > 0 {
		resp.Where("nombre LIKE ?", "%"+filtro.Nombre+"%")
	}

	if len(filtro.CodigoAcceso) > 0 {
		resp.Where("codigo_acceso = ?", filtro.CodigoAcceso).Select("comunidades.id, comunidades.nombre")
	}

	resp.Order("comunidades.created_at DESC")
	resp.Find(&comunidades)

	if resp.Error != nil {
		log := entities.Log{
			Tipo:          entities.Error,
			Mensaje:       resp.Error.Error(),
			Funcionalidad: "GetComunidadesRepository",
		}
		if err := r.utilService.CreateLogService(log); err != nil {
			mensaje := fmt.Sprintf("Crear Log: %s. %s", err.Error(), err.Error())
			logs.Error(mensaje)
		}
	} else if resp.RowsAffected <= 0 {
		erro = errors.New("no se encontraron comunidades")
	}
	return
}

func (r *comunidadRepository) PostTrayectoRepository(trayecto entities.Trayecto) error {
	return r.SqlClient.Transaction(func(tx *gorm.DB) error {

		// 1. Crear el Trayecto (cabecera)
		if err := tx.Create(&trayecto).Error; err != nil {
			logs.Info(err)
			return errors.New("error al guardar el trayecto")
		}

		// // 2. Crear el TrayectoDetalle (detalle) asociado al Trayecto
		// // Si el TrayectoDetalle está anidado dentro de Trayecto, GORM lo guardará automáticamente
		// if err := tx.Create(&trayecto.Detalles).Error; err != nil {
		// 	logs.Info(err)
		// 	return errors.New("error al guardar el detalle del trayecto")
		// }

		// // 3. Crear las rutinas asociadas al trayecto
		// // Aseguramos que las rutinas están asociadas correctamente al trayecto antes de la inserción
		// for _, rutina := range trayecto.Rutinas {
		// 	rutina.TrayectosID = trayecto.ID
		// }

		// // Insertamos las rutinas asociadas al trayecto
		// if err := tx.CreateInBatches(&trayecto.Rutinas, 10).Error; err != nil {
		// 	logs.Info(err)
		// 	return errors.New("error al guardar las rutinas asociadas al trayecto")
		// }

		// Si todo fue exitoso, confirmamos la transacción
		return nil
	})
}

func (r *comunidadRepository) GetTrayectosRepository(filtro filtros.TrayectoFiltro) (trayectos []*entities.Trayecto, totalFilas int64, erro error) {

	resp := r.SqlClient.Model(&entities.Trayecto{})

	if filtro.ID > 0 {
		resp.Where("id = ?", filtro.ID).Limit(1)
	}

	if filtro.ComunidadID > 0 {
		resp.Where("comunidades_id = ?", filtro.ComunidadID)
	}

	if filtro.CargarDetalle {
		resp.Preload("Detalles")
	}

	resp.Preload("Rutinas")
	resp.Preload("Usuario")
	resp.Preload("Estado")
	resp.Order("trayectos.created_at DESC")

	// PAGINACIÓN
	if filtro.Number > 0 && filtro.Size > 0 {
		resp.Count(&totalFilas)
		if resp.Error != nil {
			erro = fmt.Errorf(ERROR_CARGAR_TOTAL_FILAS)
		}
		offset := (filtro.Number - 1) * filtro.Size
		resp.Limit(int(filtro.Size))
		resp.Offset(int(offset))
	}

	resp.Find(&trayectos)

	if resp.Error != nil {
		log := entities.Log{
			Tipo:          entities.Error,
			Mensaje:       resp.Error.Error(),
			Funcionalidad: "GetTrayectosRepository",
		}
		if err := r.utilService.CreateLogService(log); err != nil {
			mensaje := fmt.Sprintf("Crear Log: %s. %s", err.Error(), err.Error())
			logs.Error(mensaje)
		}
	} else if resp.RowsAffected <= 0 {
		erro = errors.New("no se encontraron trayectos")
	}
	return
}

func (r *comunidadRepository) PostUsuarioRolComunidadReporitory(entity entities.UsuarioRolComunidad) (erro error) {
	result := r.SqlClient.Model(entities.UsuarioRolComunidad{}).Create(&entity)

	if result.Error != nil {
		erro = fmt.Errorf("error al registrar el miembro %d en la comunidad %d", entity.UsuariosID, entity.ComunidadesID)
		log := entities.Log{
			Tipo:          entities.Error,
			Mensaje:       result.Error.Error(),
			Funcionalidad: "PostUsuarioRolComunidadReporitory",
		}

		err := r.utilService.CreateLogService(log)
		if err != nil {
			mensaje := fmt.Sprintf("%s, %s", err.Error(), result.Error.Error())
			logs.Error(mensaje)
		}
	}
	return
}

func (r *comunidadRepository) UpdateTrayectoRepository(id uint, updateFields map[string]interface{}) (erro error) {
	if err := r.SqlClient.Model(&entities.Trayecto{}).
		Where("id = ?", id).
		Updates(updateFields).Error; err != nil {
		return err
	}
	return
}
