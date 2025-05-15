package comunidad

import (
	"errors"
	"fmt"

	"github.com/faridEmilio/api_go_viajate_corporativo/internal/database"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
	"gorm.io/gorm"
)

type ComunidadRepository interface {
	GetComunidadesRepository(request comunidaddtos.RequestComunidad) (comunidades []entities.Comunidad, total int64, erro error)
	PostComunidadRepository(comunidad entities.Comunidad) (erro error)
	UpdateComunidadRepository(comunidad entities.Comunidad) (erro error)
	PostUsuarioComunidadRepository(usuariocomunidad entities.UsuariosHasComunidades) (erro error)
	GetUsuarioComunidadRepository(request comunidaddtos.RequestAltaMiembro) (usuariocomunidad []entities.UsuariosHasComunidades, erro error)
	UpdateUsuarioComunidadRepository(usuariocomunidad entities.UsuariosHasComunidades) (erro error)
	GetTipoComunidadRepository(request comunidaddtos.RequestTipoComunidad) (tipocomunidad []entities.TipoComunidad, total int64, erro error)
	// CRUD TRAYECTO
	// PostTrayectoRepository(trayecto entities.Trayecto) error
	// GetTrayectosRepository(filtro filtros.TrayectoFiltro) (trayectos []*entities.Trayecto, totalFilas int64, erro error)
	// UpdateTrayectoRepository(id uint, updateFields map[string]interface{}) (erro error)

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

func (r *comunidadRepository) GetComunidadesRepository(request comunidaddtos.RequestComunidad) (comunidades []entities.Comunidad, total int64, erro error) {
	resp := r.SqlClient.Model(&entities.Comunidad{})

	// FILTROS
	if request.ID > 0 {
		resp.Where("comunidades.id = ?", request.ID).Limit(1)
	}

	if len(request.Nombre) > 0 {
		resp.Where("comunidades.nombre REGEXP ?", request.Nombre)
	}

	if len(request.CodigoAcceso) > 0 {
		resp.Where("comunidades.codigo_acceso = ?", request.CodigoAcceso)
	}

	resp = resp.Joins("INNER JOIN localidades l ON l.id = comunidades.localidades_id").
		Preload("Localidad", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "nombre", "provincias_id").
				Preload("Provincia", func(db *gorm.DB) *gorm.DB {
					return db.Select("id", "nombre", "paises_id").
						Preload("Pais", func(db *gorm.DB) *gorm.DB {
							return db.Select("id", "nombre")
						})
				})
		})

	resp = resp.Joins("INNER JOIN tipo_comunidad tp ON tp.id = comunidades.tipo_comunidad_id").
		Preload("TipoComunidad", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "tipo")
		})
	if request.Number > 0 && request.Size > 0 {

		resp.Count(&total)

		if resp.Error != nil {
			erro = fmt.Errorf(ERROR_CARGAR_TOTAL_FILAS)
		}

		offset := (request.Number - 1) * request.Size
		resp.Limit(int(request.Size))
		resp.Offset(int(offset))
	}

	resp.Order("comunidades.created_at DESC")
	resp.Find(&comunidades)
	if resp.Error != nil {
		erro = fmt.Errorf(ERROR_CONSULTA, erro.Error())
		return
	}

	return
}

func (r *comunidadRepository) PostComunidadRepository(comunidad entities.Comunidad) (erro error) {
	err := r.SqlClient.Create(&comunidad).Error
	if err != nil {
		erro = errors.New("no se pudo crear la comunidad")
		return
	}

	return
}

func (r *comunidadRepository) UpdateComunidadRepository(comunidad entities.Comunidad) (erro error) {
	err := r.SqlClient.Save(&comunidad).Error
	if err != nil {
		erro = errors.New("no se pudo actualizar la comunidad")
		return
	}
	return
}

func (r *comunidadRepository) GetUsuarioComunidadRepository(request comunidaddtos.RequestAltaMiembro) (usuariocomunidad []entities.UsuariosHasComunidades, erro error) {
	resp := r.SqlClient.Model(&entities.Comunidad{}).Where("comunidades_id = ? AND usuarios_id = ?", request.ComunidadId, request.UsuariosID).Find(&usuariocomunidad)
	if resp.Error != nil {
		erro = fmt.Errorf(ERROR_CONSULTA, erro.Error())
		return
	}
	return
}

func (r *comunidadRepository) PostUsuarioComunidadRepository(usuariocomunidad entities.UsuariosHasComunidades) (erro error) {
	err := r.SqlClient.Create(&usuariocomunidad).Error
	if err != nil {
		erro = errors.New("no se pudo crear la comunidad")
		return
	}

	return
}

func (r *comunidadRepository) UpdateUsuarioComunidadRepository(usuariocomunidad entities.UsuariosHasComunidades) (erro error) {
	err := r.SqlClient.Save(&usuariocomunidad).Error
	if err != nil {
		erro = errors.New("no se pudo actualizar usuario comunidades")
		return
	}
	return
}

func (r *comunidadRepository) GetTipoComunidadRepository(request comunidaddtos.RequestTipoComunidad) (tipocomunidad []entities.TipoComunidad, total int64, erro error) {
	resp := r.SqlClient.Model(&entities.TipoComunidad{})
	if request.Id > 0 {
		resp.Where("id = ?", request.Id)
	}
	if request.Number > 0 && request.Size > 0 {

		resp.Count(&total)

		if resp.Error != nil {
			erro = fmt.Errorf(ERROR_CARGAR_TOTAL_FILAS)
		}

		offset := (request.Number - 1) * request.Size
		resp.Limit(int(request.Size))
		resp.Offset(int(offset))
	}
	resp.Find(&tipocomunidad)
	if resp.Error != nil {
		erro = fmt.Errorf(ERROR_CONSULTA, erro.Error())
		return
	}
	return
}

// func (r *comunidadRepository) PostTrayectoRepository(trayecto entities.Trayecto) error {
// 	return r.SqlClient.Transaction(func(tx *gorm.DB) error {

// 		// 1. Crear el Trayecto (cabecera)
// 		if err := tx.Create(&trayecto).Error; err != nil {
// 			logs.Info(err)
// 			return errors.New("error al guardar el trayecto")
// 		}

// 		// // 2. Crear el TrayectoDetalle (detalle) asociado al Trayecto
// 		// // Si el TrayectoDetalle está anidado dentro de Trayecto, GORM lo guardará automáticamente
// 		// if err := tx.Create(&trayecto.Detalles).Error; err != nil {
// 		// 	logs.Info(err)
// 		// 	return errors.New("error al guardar el detalle del trayecto")
// 		// }

// 		// // 3. Crear las rutinas asociadas al trayecto
// 		// // Aseguramos que las rutinas están asociadas correctamente al trayecto antes de la inserción
// 		// for _, rutina := range trayecto.Rutinas {
// 		// 	rutina.TrayectosID = trayecto.ID
// 		// }

// 		// // Insertamos las rutinas asociadas al trayecto
// 		// if err := tx.CreateInBatches(&trayecto.Rutinas, 10).Error; err != nil {
// 		// 	logs.Info(err)
// 		// 	return errors.New("error al guardar las rutinas asociadas al trayecto")
// 		// }

// 		// Si todo fue exitoso, confirmamos la transacción
// 		return nil
// 	})
// }

// func (r *comunidadRepository) GetTrayectosRepository(filtro filtros.TrayectoFiltro) (trayectos []*entities.Trayecto, totalFilas int64, erro error) {

// 	resp := r.SqlClient.Model(&entities.Trayecto{})

// 	if filtro.ID > 0 {
// 		resp.Where("id = ?", filtro.ID).Limit(1)
// 	}

// 	if filtro.ComunidadID > 0 {
// 		resp.Where("comunidades_id = ?", filtro.ComunidadID)
// 	}

// 	if filtro.CargarDetalle {
// 		resp.Preload("Detalles")
// 	}

// 	resp.Preload("Rutinas")
// 	resp.Preload("Usuario")
// 	resp.Preload("Estado")
// 	resp.Order("trayectos.created_at DESC")

// 	// PAGINACIÓN
// 	if filtro.Number > 0 && filtro.Size > 0 {
// 		resp.Count(&totalFilas)
// 		if resp.Error != nil {
// 			erro = fmt.Errorf(ERROR_CARGAR_TOTAL_FILAS)
// 		}
// 		offset := (filtro.Number - 1) * filtro.Size
// 		resp.Limit(int(filtro.Size))
// 		resp.Offset(int(offset))
// 	}

// 	resp.Find(&trayectos)

// 	if resp.Error != nil {
// 		log := entities.Log{
// 			Tipo:          entities.Error,
// 			Mensaje:       resp.Error.Error(),
// 			Funcionalidad: "GetTrayectosRepository",
// 		}
// 		if err := r.utilService.CreateLogService(log); err != nil {
// 			mensaje := fmt.Sprintf("Crear Log: %s. %s", err.Error(), err.Error())
// 			logs.Error(mensaje)
// 		}
// 	} else if resp.RowsAffected <= 0 {
// 		erro = errors.New("no se encontraron trayectos")
// 	}
// 	return
// }

// func (r *comunidadRepository) UpdateTrayectoRepository(id uint, updateFields map[string]interface{}) (erro error) {
// 	if err := r.SqlClient.Model(&entities.Trayecto{}).
// 		Where("id = ?", id).
// 		Updates(updateFields).Error; err != nil {
// 		return err
// 	}
// 	return
// }
