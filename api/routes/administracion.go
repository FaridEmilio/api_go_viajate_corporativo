package routes


func AdministracionRoutes(app fiber.Router, middlewares middlewares.MiddlewareManager, administracionService administracion.AdministracionService, utilService util.UtilService, commons commons.Commons, runEndpoint util.RunEndpoint) {
	app.Get("/:comunidad_id/miembros", middlewares.ValidarPermisoComunidad("ADMINISTRADOR"), GetMiembros(administracionService))
}