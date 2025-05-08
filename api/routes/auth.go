package routes

import (
	"log"

	"github.com/faridEmilio/api_go_viajate_corporativo/api/middlewares"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/auth"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/authdtos"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router, middlewares middlewares.MiddlewareManager, authService auth.AuthService, utilService util.UtilService) {
	// app.Post("/profile-photo", middlewares.ValidarPermiso(), UploadProfilePicture(usuarioService))
	// //app.Delete("/profile-photo", middlewares.ValidarPermiso(), DeleteProfilePicture(usuarioService))
	// app.Put("/change-password", middlewares.ValidarPermiso(), ChangePassword(usuarioService))
	// app.Put("/update", middlewares.ValidarPermiso(), UpdateUsuario(usuarioService))

	/* ---------------------------- AUTH ---------------------------- */
	app.Post("/register", Register(authService))
	app.Post("/login", Login(authService))

	// app.Get("/user/:id", GetUser(authService))
	// app.Post("/refresh-token", middlewares.ValidarPermiso(), RefreshToken(authService))
	// app.Post("/verify-email", VerifyEmail(authService))

	// //Recover Password
	// app.Post("/restore-password", RestorePassword(viajateService, runEndpoint))
	// app.Post("/reset-password", ResetPassword(viajateService))
}

func Login(authService auth.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request authdtos.RequestLogin
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body"})
		}

		// Llama al servicio de autenticación
		data, err := authService.Login(request)
		if err != nil {
			log.Println("Login error:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  false,
				"message": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  true,
			"data":    data,
			"message": "Sesión iniciada con éxito",
		})
	}
}

func Register(authService auth.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var request authdtos.RequestNewUser
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body"})
		}

		user, err := authService.Register(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}

		data, err := authService.GetTokensService(user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  true,
			"data":    data,
			"message": "Registro completado con éxito",
		})
	}
}
