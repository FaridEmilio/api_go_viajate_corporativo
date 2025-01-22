package administracion

import (
	"github.com/faridEmilio/api_go_gym_manager/internal/database"
	
)

type Repository interface {
}

func NewAdministracionRepository(conn *database.MySQLClient, util util.UtilService) Repository {
	return &repository{
		SQLClient:   conn,
		utilService: util,
	}
}

type repository struct {
	SQLClient   *database.MySQLClient
	utilService util.UtilService
}

// func (r *repository) GetMiembrosRepository(filtro filtros.MiembroFiltro) (miembros []entities.Usuario, totalFilas int64, erro error) {
// 	resp := r.SQLClient.Model(entities.Usuario{})

// 	//if filtro.ComunidadID > 0 {
// 	resp = resp.Joins("INNER JOIN usuarios_roles_comunidades urc ON urc.usuarios_id = usuarios.id").
// 		Where("urc.comunidades_id = ?", filtro.ComunidadID).
// 		Preload("RolesComunidad", "comunidades_id = ?", filtro.ComunidadID)
// 	//}

// 	// Excluir al usuario que hizo la consulta
// 	if filtro.UsuarioID > 0 {
// 		resp = resp.Where("usuarios.id <> ?", filtro.UsuarioID)
// 	}

// 	if len(filtro.Apellido) > 0 {
// 		resp.Where("apellido LIKE ?", "%"+filtro.Apellido+"%")
// 	}

// 	if len(filtro.Nombre) > 0 {
// 		resp.Where("nombre LIKE ?", "%"+filtro.Nombre+"%")
// 	}

// 	if filtro.SexoMasculino {
// 		resp.Where("genero = ?", "Masculino")
// 	}

// 	if filtro.SexoFedatabase
// 		resp.Where("genero = ?", "Femenino")
// 	}

// 	if filtro.OrdenarFechaNacimiento {
// 		resp.Order("fecha_nacimiento ASC")
// 	}

// 	// Filtrar por rol: Miembro, Administrador o ambos
// 	if filtro.CargarSoloMiembros {
// 		resp = resp.Where("urc.rol = ?", "MIEMBRO")
// 	} else if filtro.CargarSoloAdministradores {
// 		resp = resp.Where("urc.rol = ?", "ADMINISTRADOR")
// 	}
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

// 	resp.Find(&miembros)

// 	if resp.Error != nil {
// 		log := entities.Log{
// 			Tipo:          entities.Error,
// 			Mensaje:       resp.Error.Error(),
// 			Funcionalidad: "GetMiembrosRepository",
// 		}
// 		if err := r.utilService.CreateLogService(log); err != nil {
// 			mensaje := fmt.Sprintf("Crear Log: %s. %s", err.Error(), err.Error())
// 			logs.Error(mensaje)
// 		}
// 	} else if resp.RowsAffected <= 0 {
// 		erro = errors.New("no se encontraron miembros")
// 	}
// 	return
// }

// func (r *repository) UpdateComunidadRepository(comunidadID uint, updateFields map[string]interface{}) (erro error) {
// 	if err := r.SQLClient.Model(&entities.Comunidad{}).
// 		Where("id = ?", comunidadID).
// 		Updates(updateFields).Error; err != nil {
// 		return err
// 	}
// 	return
// }
