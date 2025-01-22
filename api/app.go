package main

import (
	"net/http"

	"github.com/faridEmilio/api_go_gym_manager/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	//"github.com/gofiber/template/html"
)

func InicializarApp(clienteHttp *http.Client) *fiber.App {
	//Servicios comunes
	// fileRepository := commons.NewFileRepository(clienteFile)
	// commonsService := commons.NewCommons(fileRepository)
	// algoritmoVerificacionService := commons.NewAlgoritmoVerificacion()
	//middlewares := middlewares.MiddlewareManager{HTTPClient: clienteHttp}

	//utilRepository := util.NewUtilRepository(clienteSql)
	//utilService := util.NewUtilService(utilRepository)

	// //Valida si existe un correo para solicitud de nuevas cuentas si no existe lo crea.
	// utilService.FirstOrCreateConfiguracionService("EMAIL_SOLICITUD_CUENTA", "Email que recibirá la solicitud de apertura de cuenta", "developmenttelco@gmail.com")

	//ApiLink
	// apiLinkRemoteRepository := apilink.NewRemote(clienteHttp, utilService)
	// apiLinkService := apilink.NewService(apiLinkRemoteRepository)

	// auditoriaRespository := auditoria.NewAuditoriaRepository(clienteSql)
	// auditoriaService := auditoria.AuditoriaService(auditoriaRespository)

	// administracionRepository := administracion.NewRepository(clienteSql, auditoriaService, utilService)
	// administracionService := administracion.NewService(administracionRepository, apiLinkService, commonsService, utilService)

	// usuarioRepository := usuario.NewRepository(clienteSql, utilService)
	// usuarioRemoteRepository := usuario.NewRemote(clienteHttp)
	// usuarioService := usuario.NewService(usuarioRemoteRepository, usuarioRepository)

	// remoteRepository := prisma.NewRepoasitory(clienteHttp)
	// prismaRepository := prisma.NewRepository(clienteSql)
	// prismaService := prisma.NewService(remoteRepository, prismaRepository, commonsService)
	// pagoOffLineService := pagooffline.NewService(algoritmoVerificacionService)
	// cierreloteRepository := cierrelote.NewRepository(clienteSql)
	// storage := storage.NewS3Session()
	// reafileStore := cierrelote.NewStore(storage)
	// cierreloteService := cierrelote.NewService(cierreloteRepository, commonsService, utilService, reafileStore)
	// checkoutRepository := checkout.NewRepository(clienteSql, auditoriaService)
	// checkoutService := checkout.NewService(checkoutRepository, commonsService, prismaService, pagoOffLineService)

	//engine := html.New(filepath.Join(filepath.Base("."), "api", "views"), ".html")
	//engine := html.New("views", ".html")
	//engine.Delims("${", "}")
	app := fiber.New(fiber.Config{
		//Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var msg string
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Message
			}

			if msg == "" {
				msg = "No se pudo procesar el llamado a la api: " + err.Error()
			}

			_ = ctx.Status(code).JSON(internalError{
				Message: msg,
			})

			return nil
		},
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Corrientes Telecomunicaciones Api Servicio de Pasarela de Pagos"))
	})

	api := app.Group("/api/v1")
	routes.PruebaRoutes(api)
	//aca mando el parametro
	routes.RegistroRoutes(api)

	// checkout := app.Group("/checkout")
	// routes.CheckoutRoutes(checkout, checkoutService)

	//pagooffline := app.Group("/pagooffline")
	//routes.PrismaRoutes(pagooffline, pagoOffLineService)

	// prisma := app.Group("/prisma")
	// routes.PrismaRoutes(prisma, prismaService)

	// cierrelote := app.Group("/cierrelote")
	// routes.CierreLoteRoutes(cierrelote, cierreloteService, administracionService)

	// administracion := app.Group("/administracion")
	// routes.AdministracionRoutes(administracion, middlewares, administracionService, utilService)

	// usuario := app.Group("/usuario")
	// routes.UsuarioRoutes(usuario, middlewares, usuarioService)

	// //Procesos en segundo plano
	// background.BackgroudServices(administracionService, cierreloteService, utilService)

	//app.Static("/", "./views")
	//app.Static("/", filepath.Join(filepath.Base("."), "api", "views"))

	return app
}

func main() {

	//HTTPClient.Timeout = time.Second * 120 //Todo validar si este tiempo está bien

	app := InicializarApp(http.DefaultClient)
	// el puerto puede que se necesite cambiar
	_ = app.Listen(":3300")
}

type internalError struct {
	Message string `json:"message"`
}
