package routes

import (
	"github.com/gofiber/fiber/v2"
)

func PruebaRoutes(app fiber.Router) {
	app.Get("/prueba", getPruebas())

}

func getPruebas() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{
			"message": "Hola mundo",
		})
	}
}
