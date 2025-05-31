package routes

import (
	"fmt"

	"github.com/faridEmilio/api_go_viajate_corporativo/api/middlewares"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/comunidad"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/authdtos"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/comunidaddtos"
	"github.com/gofiber/fiber/v2"
)

func ComunidadRoutes(app fiber.Router, middlewares middlewares.MiddlewareManager, comunidadService comunidad.ComunidadService, utilService util.UtilService) {
	//CRUD COMUNIDAD
	app.Get("/comunidades", middlewares.ValidarPermiso("admin.comunidad"), GetComunidades(comunidadService))
	app.Post("/comunidad", middlewares.ValidarPermiso("admin.comunidad"), PostComunidad(comunidadService))
	app.Post("/update-comunidad", middlewares.ValidarPermiso("admin.comunidad"), PutComunidad(comunidadService))
	app.Post("/registrar-usuario-comunidad", middlewares.ValidarPermiso("admin.comunidad"), PostUsuarioComunidad(comunidadService))

	app.Get("/tipo-comunidad", middlewares.ValidarPermiso("admin.comunidad"), GetTipoComunidad(comunidadService))

	// CRUD TRAYECTO
	app.Post("/:comunidad/route", middlewares.ValidarPermiso("crud.route"), PostRoute(comunidadService))
	// app.Get("/:comunidad_id/routes", GetRoutes(comunidadService))

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
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}
		condicion := true
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println("Error al leer el form-data")
			condicion = false
		}
		if condicion {
			files := form.File["foto_perfil"]
			if len(files) > 0 {
				file := files[0]
				urlFoto, err := comunidadService.UploadImageToFirebase(file)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "Error subiendo la imagen",
					})
				}
				request.FotoPerfil = urlFoto
			}
		}

		err = comunidadService.PostComunidadService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": "comunidad registrada con exito",
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
		var request comunidaddtos.RequestAltaMiembro
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar la solicitud",
			})
		}

		nombre, err := comunidadService.PostUsuarioComunidadService(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}
		message := "Usuario registrado en la " + nombre + " con exito"
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": message,
		})
	}
}

func PutUsuarioComunidad(comunidadService comunidad.ComunidadService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request comunidaddtos.RequestAltaMiembro

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

// func GetRoutes(comunidadService comunidad.ComunidadService) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// Obtener el comunidad_id desde el path parameter
// 		comunidadID := c.Params("comunidad_id")
// 		if comunidadID == "" {
// 			return fiber.NewError(fiber.StatusBadRequest, "Debe proporcionar un ID de comunidad válido")
// 		}
// 		comunidad_id, err := strconv.Atoi(comunidadID)
// 		if err != nil || comunidad_id <= 0 {
// 			return fiber.NewError(fiber.StatusInternalServerError, "Error al obtener la comunidad solicitada")
// 		}

// 		var request filtros.TrayectoFiltro
// 		err = c.QueryParser(&request)
// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Error en los parámetros enviados",
// 			})
// 		}

// 		// Asigno el filtro de comunidad para obtener trayectos
// 		request.ComunidadID = uint(comunidad_id)
// 		response, err := comunidadService.GetTrayectosService(request)
// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
// 				"status":  false,
// 				"message": "No se pudo obtener mis comunidades. " + err.Error(),
// 			})
// 		}

// 		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
// 			"status":  true,
// 			"data":    response,
// 			"message": "Operación de consulta de rutinas exitosa",
// 		})
// 	}
// }

// func PostUsuarioComunidad(comunidadService comunidad.ComunidadService) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		user := c.Locals("user").(viajatedtos.ResponseUsuario)
// 		if user.ID == 0 {
// 			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
// 				"status":  false,
// 				"message": "No se pudo obtener el id del usuario loggeado",
// 			})
// 		}

// 		var request comunidaddtos.RequestAltaMiembro
// 		err := c.BodyParser(&request)
// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Error al analizar la solicitud",
// 			})
// 		}

// 		request.UsuariosID = user.ID // Usuario que se quiere matricular
// 		comunidad, err := comunidadService.PostAltaMiembroService(request)
// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 		}

// 		message := fmt.Sprintf("¡Todo listo! Tu registro en la comunidad %s ha sido completado con éxito.", comunidad)
// 		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
// 			"status":  true,
// 			"message": message,
// 		})
// 	}
// }

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
