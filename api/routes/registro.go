package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegistroRoutes(app fiber.Router) {
	app.Get("/registro", getRegistro())
}

func getRegistro() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{
			"usuario": "Emilio Barrios",
		})
	}
}
