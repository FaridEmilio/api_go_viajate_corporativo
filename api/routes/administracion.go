package routes

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/api/middlewares"
	"github.com/faridEmilio/api_go_viajate_corporativo/internal/logs"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/administracion"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/administraciondtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/authdtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/administracion"
	UsuarioFiltros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/usuarios"
	"github.com/gofiber/fiber/v2"
)

func AdministracionRoutes(app fiber.Router, middlewares middlewares.MiddlewareManager, administracionService administracion.AdministracionService, utilService util.UtilService) {
	app.Get("/paises", GetPaises(administracionService))

	// expulsar miembro---- eliminar registro tabla comunidad-has-user
	app.Put("/expulsar-miembro", PutUsuaroHasComunidad(administracionService))

	// listar todos los miembros de una comunidad
	app.Get("/:comunidad_id/members", middlewares.ValidarPermiso("admin.comunidad"), GetComunidadMembers(administracionService))
	app.Get("/user", middlewares.ValidarPermiso("admin.user.search"), GetUser(administracionService))

	// crud sedes con el post con address
	app.Get("/sedes", GetSedes(administracionService))
	app.Post("/sede", CreateSede(administracionService))
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

func PutUsuaroHasComunidad(administracionService administracion.AdministracionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request comunidaddtos.RequestMiembro
		err := c.BodyParser(&request)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		err = administracionService.PutUsuarioHasComunidadService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "Error al eliminar el usuario de la comunidad. " + err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": "Usuario eliminado de la comunidad con exito",
		})
	}
}

func GetComunidadMembers(administracionService administracion.AdministracionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request filtros.MiembroFiltro
		err := c.QueryParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error en los parámetros enviados",
			})
		}

		// Filtros fijos para este endpoint
		comunidadID := uint(c.Locals("comunidadID").(int))
		userID := c.Locals("user").(authdtos.ResponseUsuario).ID
		request.ComunidadID = comunidadID
		request.AdministradorID = userID

		response, err := administracionService.GetMiembrosService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "¡Ups! No encontramos miembros en esta comunidad",
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Miembros obtenidos con éxito",
		})
	}
}

func GetSedes(administracionService administracion.AdministracionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		comunidadID := c.Locals("comunidad").(comunidaddtos.RequestComunidad).ID
		response, err := administracionService.GetSedesService(comunidadID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "¡Ups! No encontramos las sedes.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Sedes obtenidos con éxito",
		})
	}
}

func CreateSede(administracionService administracion.AdministracionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request administraciondtos.RequestCreateSede
		err := c.BodyParser(&request)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		err = administracionService.CreateSedeService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "Error al crear la sede. " + err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": "Sede creada con exito",
		})
	}
}

func PutSede(administracionService administracion.AdministracionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request administraciondtos.RequestCreateSede
		err := c.BodyParser(&request)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		err = administracionService.UpdateSedeActivaService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "Error al eliminar el usuario de la comunidad. " + err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": "Usuario eliminado de la comunidad con exito",
		})
	}
}

func GetUser(administracionService administracion.AdministracionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request UsuarioFiltros.UsuarioFiltro
		err := c.QueryParser(&request)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		response, err := administracionService.GetUsuarioService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "Lo siento, no hemos encontrado usuarios que coincidan con tu búsqueda",
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Usuario obtenidos con éxito",
		})
	}
}
