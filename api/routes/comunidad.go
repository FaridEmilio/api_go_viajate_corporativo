package routes

import (
	"fmt"

	"github.com/faridEmilio/api_go_viajate_corporativo/api/middlewares"
	"github.com/faridEmilio/api_go_viajate_corporativo/internal/logs"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/comunidad"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/authdtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/comunidad"
	"github.com/gofiber/fiber/v2"
)

func ComunidadRoutes(app fiber.Router, middlewares middlewares.MiddlewareManager, comunidadService comunidad.ComunidadService, utilService util.UtilService) {
	//CRUD COMUNIDAD
	// ** SUPERADMIN crean y ven las comunidades
	app.Post("/comunidad", middlewares.ValidarPermiso("create.comunidad"), PostComunidad(comunidadService))
	app.Get("/comunidades", middlewares.ValidarPermiso("show.comunidades"), GetComunidades(comunidadService))

	// ** ADMIN administran  comunidades
	app.Post("/update-comunidad", middlewares.ValidarPermiso("admin.comunidad"), PutComunidad(comunidadService))
	app.Post("/miembro", middlewares.ValidarPermiso("admin.comunidad"), PostUsuarioComunidad(comunidadService))

	app.Get("/tipo-comunidad", middlewares.ValidarPermiso("admin.comunidad"), GetTipoComunidad(comunidadService))

	// CRUD TRAYECTO
	app.Post("/:comunidad_id/route", middlewares.ValidarPermiso("crud.route"), PostRoute(comunidadService))
	app.Get("/:comunidad_id/routes", middlewares.ValidarPermiso("crud.route"), GetRoutes(comunidadService))

	// CRUD VEHICULO
	app.Post("/vehiculo", middlewares.ValidarPermiso("crud.vehiculo"), PostVehiculo(comunidadService))
	app.Get("/marcas", middlewares.ValidarPermiso("crud.vehiculo"), GetMarcas(comunidadService))
	app.Get("/mis-vehiculos", middlewares.ValidarPermiso("crud.vehiculo"), GetMisVehiculos(comunidadService))
}

func GetComunidades(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request comunidaddtos.RequestComunidad
		err := c.QueryParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error en los parámetros enviados",
			})
		}

		response, err := comunidadService.GetComunidadesService(request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "No se pudo obtener las comunidades. " + err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Operación de consulta de comunidades exitosa.",
		})
	}
}

func PostComunidad(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request comunidaddtos.RequestComunidad
		if err := c.BodyParser(&request); err != nil {
			logs.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		response, err := comunidadService.PostComunidadService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Comunidad creada con éxito",
		})
	}
}

func PutComunidad(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request comunidaddtos.RequestComunidad

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}
		err = comunidadService.PutComunidadService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": "comunidad actualizada con exito",
		})
	}
}

func PostUsuarioComunidad(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request comunidaddtos.RequestMiembro
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		comunidad, err := comunidadService.PostUsuarioComunidadService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": fmt.Sprintf("¡Todo listo! Tu registro en la comunidad %s ha sido completado con éxito", comunidad),
		})
	}
}

func PutUsuarioComunidad(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request comunidaddtos.RequestMiembro

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}
		err = comunidadService.PutUsuarioComunidadService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": "usuario comunidad actualizada con exito",
		})
	}
}

func GetTipoComunidad(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request comunidaddtos.RequestTipoComunidad
		err := c.QueryParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error en los parámetros enviados",
			})
		}

		response, err := comunidadService.GetTipoComunidadService(request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "No se pudo obtener las comunidades. " + err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Operación de consulta de tipo comunidades exitosa.",
		})
	}
}

// func GetMisComunidades(comunidadService comunidad.ComunidadService) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		user := c.Locals("user").(viajatedtos.ResponseUsuario)
// 		if user.ID == 0 {
// 			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
// 				"status":  false,
// 				"message": "No se pudo obtener el id del usuario loggeado",
// 			})
// 		}

// 		filtro := filtros.ComunidadFiltro{
// 			UsuarioID: user.ID,
// 		}

// 		response, err := comunidadService.GetMisComunidadesService(filtro)

// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
// 				"status":  false,
// 				"message": "No se pudo obtener mis comunidades. " + err.Error(),
// 			})
// 		}
// 		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
// 			"status":  true,
// 			"data":    response,
// 			"message": "Operación de consulta de comunidades exitosa.",
// 		})
// 	}
// }

// CRUD TRAYECTO
func PostRoute(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request comunidaddtos.RequestTrayecto
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		err := comunidadService.PostTrayectoService(request)
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
		var request filtros.TrayectoFiltro
		err := c.QueryParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error en los parámetros enviados",
			})
		}

		// Asigno el filtro de comunidad para obtener trayectos
		comunidadID := uint(c.Locals("comunidadID").(int))
		request.ComunidadID = comunidadID
		response, err := comunidadService.GetTrayectosService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "Parece que no hay trayectos disponibles en este momento. Prueba variar las fechas",
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Operación de consulta de rutinas exitosa",
		})
	}
}

func PostVehiculo(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request comunidaddtos.RequestVehiculo
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		userID := c.Locals("user").(authdtos.ResponseUsuario).ID
		request.UsuariosID = userID
		resp, err := comunidadService.PostVehiculoService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "Error al guardar vehículo. " + err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    resp,
			"message": "Vehículo guardado con éxito",
		})
	}
}

func GetMisVehiculos(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user").(authdtos.ResponseUsuario).ID
		response, err := comunidadService.GetMisVehiculosService(userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "¡Ups! No encontramos vehículos en tu lista. Añade tu primer vehículo para comenzar",
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Vehiculos obtenidos con éxito",
		})
	}
}

func GetMarcas(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		response, err := comunidadService.GetMarcasService()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": "Error al obtener marcas",
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"data":    response,
			"message": "Marcas obtenidas con éxito",
		})
	}
}
