package routes

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/api/middlewares"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/auth"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router, middlewares middlewares.MiddlewareManager, authService auth.AuthService, utilService util.UtilService) {
	// app.Post("/profile-photo", middlewares.ValidarPermiso(), UploadProfilePicture(usuarioService))
	// //app.Delete("/profile-photo", middlewares.ValidarPermiso(), DeleteProfilePicture(usuarioService))
	// app.Put("/change-password", middlewares.ValidarPermiso(), ChangePassword(usuarioService))
	// app.Put("/update", middlewares.ValidarPermiso(), UpdateUsuario(usuarioService))

	// ******** AUTH

}
