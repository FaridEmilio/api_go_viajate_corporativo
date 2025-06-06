package administracion

import (
	"errors"

	"github.com/faridEmilio/api_go_viajate_corporativo/internal/database"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/administraciondtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/entities"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/administracion"
	usuarioFiltros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/usuarios"
	"gorm.io/gorm"
)

type Repository interface {
	GetPaisesRepository(filtro filtros.PaisFiltro) (paises []entities.Pais, erro error)
	UpdateUsuarioHasComunidadRepository(request comunidaddtos.RequestMiembro) (erro error)
	GetMiembrosRepository(filtro filtros.MiembroFiltro) (miembros []entities.UsuariosHasComunidades, err error)
	GetSedesRepository(comunidadID uint) (sedes []entities.Sede, err error)
	CreateSedeRepository(request administraciondtos.RequestCreateSede) (erro error)
	UpdateSedeActivaRepository(sedeID uint, activo bool) (erro error)

	GetUsuariosRepository(filtro usuarioFiltros.UsuarioFiltro) (usuarios []entities.Usuario, totalFilas int64, erro error)
	GetUsuarioRepository(filtro usuarioFiltros.UsuarioFiltro) (user entities.Usuario, erro error)
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

func (r *repository) GetPaisesRepository(filtro filtros.PaisFiltro) (paises []entities.Pais, erro error) {
	result := r.SQLClient.
		Model(&entities.Pais{}).
		Select("id", "nombre").
		Preload("Provincias", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "nombre", "paises_id").
				Preload("Localidades", func(db *gorm.DB) *gorm.DB {
					return db.Select("id", "nombre", "provincias_id")
				})
		}).
		Find(&paises)

	if result.Error != nil {
		return nil, errors.New("No se encontraron países")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("No se encontraron países")
	}

	return paises, nil
}

func (r *repository) UpdateUsuarioHasComunidadRepository(request comunidaddtos.RequestMiembro) error {
	entity := entities.UsuariosHasComunidades{
		UsuariosID:    request.UsuariosID,
		ComunidadesID: request.ComunidadID,
	}

	result := r.SQLClient.Model(&entity).
		Where("usuarios_id = ? AND comunidades_id = ?", request.UsuariosID, request.ComunidadID).
		Update("activo", false)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		erro := errors.New("no se encontró relación usuario-comunidad para actualizar")
		return erro
	}

	return nil
}

func (r *repository) GetMiembrosRepository(filtro filtros.MiembroFiltro) (miembros []entities.UsuariosHasComunidades, err error) {
	resp := r.SQLClient.Model(&entities.UsuariosHasComunidades{}).Where("comunidades_id", filtro.ComunidadID)

	if filtro.Activos {
		resp.Where("activo", true)
	}

	if filtro.Expulsados {
		resp.Where("activo", false)
	}

	//if filtro.adm

	resp.Preload("Usuario")
	resp.Find(&miembros)
	return
}

func (r *repository) GetSedesRepository(comunidadID uint) (sedes []entities.Sede, err error) {
	err = r.SQLClient.Where("comunidades_id = ? AND active = ?", comunidadID, true).
		Preload("Address").
		Find(&sedes).Error

	return sedes, err
}

func (r *repository) CreateSedeRepository(request administraciondtos.RequestCreateSede) error {
	sede, address := request.ToEntity()

	return r.SQLClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&address).Error; err != nil {
			return err
		}
		sede.AddressesID = address.ID
		if err := tx.Create(&sede).Error; err != nil {
			return err
		}
		return nil
	})
}
func (r *repository) UpdateSedeActivaRepository(sedeID uint, activo bool) (erro error) {
	return r.SQLClient.Model(&entities.Sede{}).
		Where("id = ?", sedeID).
		Update("activo", activo).Error
}

func (r *repository) GetUsuariosRepository(filtro usuarioFiltros.UsuarioFiltro) (usuarios []entities.Usuario, totalFilas int64, erro error) {
	resp := r.SQLClient.Model(entities.Usuario{}).
		Joins("JOIN usuarios_has_comunidades uhc ON uhc.usuarios_id = usuarios.id").
		Where("uhc.comunidades_id = ?", filtro.ComunidadID)

	// Excluir al usuario que hizo la consulta
	//if filtro.UsuarioID > 0 {
	//	resp = resp.Where("usuarios.id <> ?", filtro.UsuarioID)
	//}
	//
	//if len(filtro.Apellido) > 0 {
	//	resp.Where("apellido LIKE ?", "%"+filtro.Apellido+"%")
	//}
	//
	//if len(filtro.Nombre) > 0 {
	//	resp.Where("nombre LIKE ?", "%"+filtro.Nombre+"%")
	//}
	//
	//// Filtrar por rol: Miembro, Administrador o ambos
	//if filtro.SoloAdministradores {
	//	resp = resp.Where("roles_id = ?", "MIEMBRO")
	//} else if filtro.SoloAdministradores {
	//	resp = resp.Where("roles_id = ?", "ADMINISTRADOR")
	//}
	//// PAGINACIÓN
	//if filtro.Number > 0 && filtro.Size > 0 {
	//	resp.Count(&totalFilas)
	//	if resp.Error != nil {
	//		erro = fmt.Errorf(ERROR_CARGAR_TOTAL_FILAS)
	//	}
	//	offset := (filtro.Number - 1) * filtro.Size
	//	resp.Limit(int(filtro.Size))
	//	resp.Offset(int(offset))
	//}
	resp.Find(&usuarios)
	if resp.Error != nil {

	} else if resp.RowsAffected <= 0 {
		erro = errors.New("no se encontraron miembros")
	}
	return
}

// func (r *repository) UpdateComunidadRepository(comunidadID uint, updateFields map[string]interface{}) (erro error) {
// 	if err := r.SQLClient.Model(&entities.Comunidad{}).
// 		Where("id = ?", comunidadID).
// 		Updates(updateFields).Error; err != nil {
// 		return err
// 	}
// 	return
// }

func (r *repository) GetUsuarioRepository(filtro usuarioFiltros.UsuarioFiltro) (user entities.Usuario, erro error) {
	resp := r.SQLClient.Model(entities.Usuario{})
	// Si se especifican campos, usar Select para elegirlos
	if len(filtro.SelectFields) > 0 {
		resp = resp.Select(filtro.SelectFields)
	}

	if len(filtro.Email) > 0 {
		resp.Where("email = ?", filtro.Email)
	}

	if filtro.ID > 0 {
		resp.Where("id = ?", filtro.ID)
	}

	// Verificar si hay cláusulas de filtro (como WHERE) antes de ejecutar la consulta
	if len(resp.Statement.Clauses) == 0 {
		erro = errors.New("se requiere al menos un filtro para realizar la consulta")
		return
	}

	resp.First(&user)
	if resp.Error != nil {
		erro = errors.New("user not found")
		return
	}
	return
}
