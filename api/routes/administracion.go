package routes

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/api/middlewares"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/commons"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/administracion"
	"github.com/gofiber/fiber/v2"
)

func AdministracionRoutes(app fiber.Router, middlewares middlewares.MiddlewareManager, administracionService administracion.AdministracionService, utilService util.UtilService, commons commons.Commons, runEndpoint util.RunEndpoint) {
	app.Get("/:comunidad_id/miembros", middlewares.ValidarPermisoComunidad("ADMINISTRADOR"), GetMiembros(administracionService))
}
