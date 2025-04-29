package routes

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/api/middlewares"
	"github.com/faridEmilio/api_go_viajate_corporativo/internal/logs"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/administracion"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/administracion"
	"github.com/gofiber/fiber/v2"
)

func AdministracionRoutes(app fiber.Router, middlewares middlewares.MiddlewareManager, administracionService administracion.AdministracionService, utilService util.UtilService) {
	app.Get("/paises", GetPaises(administracionService))
}

func GetPaises(administracionService administracion.AdministracionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request filtros.PaisFiltro
		err := c.QueryParser(&request)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		response, err := administracionService.GetPaisesService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "Error al obtener paises. " + err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Paises obtenidos con éxito",
		})
	}
}
