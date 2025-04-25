package routes

import (
	"fmt"
	"strconv"

	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/api/middlewares"
	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/internal/logs"
	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/pkg/commons"
	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/pkg/domains/comunidad"
	"github.com/gofiber/fiber"
)

func ComunidadRoutes(app fiber.Router, middlewares middlewares.MiddlewareManager, comunidadService comunidad.ComunidadService, utilService util.UtilService, commons commons.Commons, runEndpoint util.RunEndpoint) {
	app.Post("/miembro", middlewares.ValidarPermiso(), PostUsuarioComunidad(comunidadService))

	app.Get("/:comunidad_id/mis-comunidades", middlewares.ValidarPermiso(), GetMisComunidades(comunidadService))

	// CRUD TRAYECTO
	app.Post("/:comunidad_id/new-route", middlewares.ValidarPermisoComunidad("MIEMBRO"), PostRoute(comunidadService))
	app.Get("/:comunidad_id/routes", middlewares.ValidarPermisoComunidad("MIEMBRO"), GetRoutes(comunidadService))
}

func GetMisComunidades(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(viajatedtos.ResponseUsuario)
		if user.ID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "No se pudo obtener el id del usuario loggeado",
			})
		}

		filtro := filtros.ComunidadFiltro{
			UsuarioID: user.ID,
		}

		response, err := comunidadService.GetMisComunidadesService(filtro)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "No se pudo obtener mis comunidades. " + err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Operación de consulta de comunidades exitosa.",
		})
	}
}

// CRUD TRAYECTO
func PostRoute(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(viajatedtos.ResponseUsuario)
		if user.ID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "No se pudo obtener el id del usuario loggeado",
			})
		}

		// Obtener el comunidad_id desde el path parameter
		comunidadID := c.Params("comunidad_id")
		if comunidadID == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Debe proporcionar un ID de comunidad válido")
		}
		comunidad_id, err := strconv.Atoi(comunidadID)
		if err != nil || comunidad_id <= 0 {
			return fiber.NewError(fiber.StatusInternalServerError, "Error al obtener la comunidad solicitada")
		}

		var request comunidaddtos.RequestTrayecto
		err = c.BodyParser(&request)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		// Obtengo el ID del usuario creador del trayecto
		request.UsuariosID = user.ID
		request.ComunidadesID = uint(comunidad_id)
		err = comunidadService.PostTrayectoService(request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": "Trayecto creado con éxito",
		})
	}
}

func GetRoutes(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Obtener el comunidad_id desde el path parameter
		comunidadID := c.Params("comunidad_id")
		if comunidadID == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Debe proporcionar un ID de comunidad válido")
		}
		comunidad_id, err := strconv.Atoi(comunidadID)
		if err != nil || comunidad_id <= 0 {
			return fiber.NewError(fiber.StatusInternalServerError, "Error al obtener la comunidad solicitada")
		}

		var request filtros.TrayectoFiltro
		err = c.QueryParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error en los parámetros enviados",
			})
		}

		// Asigno el filtro de comunidad para obtener trayectos
		request.ComunidadID = uint(comunidad_id)
		response, err := comunidadService.GetTrayectosService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "No se pudo obtener mis comunidades. " + err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Operación de consulta de rutinas exitosa",
		})
	}
}

func PostUsuarioComunidad(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(viajatedtos.ResponseUsuario)
		if user.ID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "No se pudo obtener el id del usuario loggeado",
			})
		}

		var request comunidaddtos.RequestAltaMiembro
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		request.UsuariosID = user.ID // Usuario que se quiere matricular
		comunidad, err := comunidadService.PostAltaMiembroService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}

		message := fmt.Sprintf("¡Todo listo! Tu registro en la comunidad %s ha sido completado con éxito.", comunidad)
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": message,
		})
	}
}
