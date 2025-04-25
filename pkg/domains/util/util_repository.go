package util

import (
	"fmt"

	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/internal/database"
	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/internal/logs"
	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/pkg/entities"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo_corporativo/pkg/filtros/utils"
)

type UtilRepository interface {
	CreateNotificacion(notificacion entities.Notificacione) error
	CreateLog(log entities.Log) (erro error)
	GetConfiguracion(filtro filtros.ConfiguracionFiltro) (configuracion entities.Configuracione, erro error)
	CreateConfiguracion(config entities.Configuracione) (id uint, erro error)
}

func NewUtilRepository(conn *database.MySQLClient) UtilRepository {
	return &utilRepository{
		SqlClient: conn,
	}
}

type utilRepository struct {
	SqlClient *database.MySQLClient
}

func (r *utilRepository) CreateNotificacion(notificacion entities.Notificacione) error {

	resp := r.SqlClient.Omit("id").Create(&notificacion)

	if resp.Error != nil {

		erro := fmt.Errorf(ERROR_CREAR_NOTIFICACION)

		log := entities.Log{
			Tipo:          entities.Error,
			Mensaje:       resp.Error.Error(),
			Funcionalidad: "CreateNotificacion",
		}

		err := r.CreateLog(log)

		if err != nil {
			mensaje := fmt.Sprintf("Crear Log: %s. %s", err.Error(), resp.Error.Error())
			logs.Error(mensaje)
		}
		return erro

	}
	return nil
}

func (r *utilRepository) CreateLog(log entities.Log) (erro error) {

	resp := r.SqlClient.Omit("id").Create(&log)

	if resp.Error != nil {
		return fmt.Errorf("error al crear log %s", resp.Error.Error())
	}

	return nil
}

func (r *utilRepository) GetConfiguracion(filtro filtros.ConfiguracionFiltro) (configuracion entities.Configuracione, erro error) {

	resp := r.SqlClient.Model(entities.Configuracione{})

	if len(filtro.Nombre) > 0 {

		resp.Where("nombre", filtro.Nombre)
	}

	resp.Find(&configuracion)

	if resp.Error != nil {

		erro = fmt.Errorf(ERROR_CONFIGURACIONES)

		log := entities.Log{
			Tipo:          entities.Error,
			Mensaje:       resp.Error.Error(),
			Funcionalidad: "GetConfiguracion",
		}

		err := r.CreateLog(log)

		if err != nil {
			mensaje := fmt.Sprintf("Crear Log: %s. %s", err.Error(), resp.Error.Error())
			logs.Error(mensaje)
		}
	}
	return
}

func (r *utilRepository) CreateConfiguracion(config entities.Configuracione) (id uint, erro error) {

	result := r.SqlClient.Omit("id").Create(&config)

	if result.Error != nil {
		erro = fmt.Errorf(ERROR_CREAR_CONFIGURACIONES)
		log := entities.Log{
			Tipo:          entities.Error,
			Mensaje:       result.Error.Error(),
			Funcionalidad: "CreateConfiguracion",
		}

		err := r.CreateLog(log)

		if err != nil {
			mensaje := fmt.Sprintf("%s, %s", err.Error(), result.Error.Error())
			logs.Error(mensaje)
		}

		return
	}

	id = config.ID

	return
}
