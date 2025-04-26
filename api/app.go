package main

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/internal/database"
	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/pkg/commons"
	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/pkg/domains/administracion"
	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/pkg/domains/comunidad"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm/logger"
)

func InicializarApp(clienteHttp *http.Client, clienteSql *database.MySQLClient, clienteFile *os.File) *fiber.App {
	//Servicios comunes
	fileRepository := commons.NewFileRepository(clienteFile)
	commonsService := commons.NewCommons(fileRepository)

	// Expo
	expoClient := expo.NewExpoClient()

	//Util
	runEndpoint := util.NewRunEndpoint(clienteHttp)
	utilRepository := util.NewUtilRepository(clienteSql)
	utilService := util.NewUtilService(utilRepository, runEndpoint)

	// Firebase Client
	firebaseClient := store.NewFirebaseClient()
	firebaseRemoteRepository := storage.NewFirebaseRemoteRepository(firebaseClient)

	// Viajate
	// REPOSITORIOS
	viajateRepository := viajate.NewViajateRepository(clienteSql, utilService)
	comunidadRepository := comunidad.NewComunidadRepository(clienteSql, utilService)
	usuarioRepository := usuario.NewUsuarioRepository(clienteSql, utilService)
	administracionRepository := administracion.NewAdministracionRepository(clienteSql, utilService)

	// SERVICIOS
	comunidadService := comunidad.NewComunidadService(comunidadRepository, utilService, commonsService)
	viajateService := viajate.NewViajateService(viajateRepository, utilService, commonsService, comunidadService, usuarioRepository)
	usuarioService := usuario.NewUsuarioService(usuarioRepository, utilService, commonsService, firebaseRemoteRepository)
	administracionService := administracion.NewAdministracionService(administracionRepository, utilService, commonsService, firebaseRemoteRepository)

	// MIDDLEWARES
	middlewares := middlewares.NewMiddlewareManager(clienteHttp, viajateService, comunidadService)
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
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://www.viajate.com.ar, http://127.0.0.1:3300, http://localhost:3000, http://localhost:8081",
		AllowHeaders:     "Content-Type, Authorization, Accept, Cookie",
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))
	app.Options("/*", func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", c.Get("Origin"))
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
		c.Set("Access-Control-Allow-Credentials", "true")
		return c.SendStatus(fiber.StatusNoContent)
	})
	// app.Use(func(ctx *fiber.Ctx) error {
	// 	config := cors.Config{
	//AllowCredentials: true,
	//AllowHeaders:     "Content-Type, Authorization",
	// }

	// if ctx.Method() == "GET" {
	// 	config.AllowOrigins = "*"
	// 	config.AllowMethods = "GET"
	// }else {
	//config.AllowOrigins = "https://viajate.com.ar, http://127.0.0.1:80"
	// 	//config.AllowMethods = "POST, PUT, DELETE"
	// }

	// 	cors.New(config)(ctx)

	// 	return ctx.Next()
	// })

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Viajate Api"))
	})

	viajateRoutes := app.Group("/api")
	routes.ViajateRoutes(viajateRoutes, middlewares, viajateService, utilService, commonsService, runEndpoint, comunidadService)

	comunidadRoutes := app.Group("/api/comunidad")
	routes.ComunidadRoutes(comunidadRoutes, middlewares, comunidadService, utilService, commonsService, runEndpoint)

	usuarioRoutes := app.Group("/api/usuario")
	routes.UsuarioRoutes(usuarioRoutes, middlewares, usuarioService, utilService, commonsService, runEndpoint)

	administracionRoutes := app.Group("/api/administracion")
	routes.AdministracionRoutes(administracionRoutes, middlewares, administracionService, utilService, commonsService, runEndpoint)
	//app.Static("/", "./views")
	//app.Static("/", filepath.Join(filepath.Base("."), "api", "views"))

	return app
}

func main() {
	var HTTPTransport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     false, // <- this is my adjustment
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	var HTTPClient = &http.Client{
		Transport: HTTPTransport,
	}

	//HTTPClient.Timeout = time.Second * 120 //Todo validar si este tiempo está bien
	clienteSQL := database.NewMySQLClient()
	osFile := os.File{}

	app := InicializarApp(HTTPClient, clienteSQL, &osFile)
	// el puerto puede que se necesite cambiar
	_ = app.Listen(":3300")
}

type internalError struct {
	Message string `json:"message"`
}
