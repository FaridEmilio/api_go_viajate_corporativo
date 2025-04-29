package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/faridEmilio/api_go_viajate_corporativo/api/middlewares"
	"github.com/faridEmilio/api_go_viajate_corporativo/api/routes"
	"github.com/faridEmilio/api_go_viajate_corporativo/internal/database"
	"github.com/faridEmilio/api_go_viajate_corporativo/internal/store"
	"github.com/joho/godotenv"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/administracion"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/auth"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/comunidad"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InicializarApp(clienteHttp *http.Client, clienteSql *database.MySQLClient, clienteFile *os.File) *fiber.App {
	//Util
	utilRepository := util.NewUtilRepository(clienteSql)
	utilService := util.NewUtilService(utilRepository)

	// Firebase Client
	firebaseClient := store.NewFirebaseClient()
	firebaseRemoteRepository := storage.NewFirebaseRemoteRepository(firebaseClient)

	// REPOSITORIOS
	comunidadRepository := comunidad.NewComunidadRepository(clienteSql, utilService)
	authRepository := auth.NewAuthRepository(clienteSql, utilService)
	administracionRepository := administracion.NewAdministracionRepository(clienteSql, utilService)

	// SERVICIOS
	comunidadService := comunidad.NewComunidadService(comunidadRepository, utilService, firebaseRemoteRepository)
	authService := auth.NewAuthService(authRepository, utilService)
	administracionService := administracion.NewAdministracionService(administracionRepository, utilService)

	// MIDDLEWARES
	middlewares := middlewares.NewMiddlewareManager(authService)
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
	// Main API Group with versioning
	api := app.Group("/api/" + os.Getenv("API_VERSION"))

	// Subgroups under api/v1
	comunidadRoutes := api.Group("/comunidad")
	routes.ComunidadRoutes(comunidadRoutes, middlewares, comunidadService, utilService)

	authRoutes := api.Group("/auth")
	routes.AuthRoutes(authRoutes, middlewares, authService, utilService)

	administracionRoutes := api.Group("/administracion")
	routes.AdministracionRoutes(administracionRoutes, middlewares, administracionService, utilService)

	//app.Static("/", "./views")
	//app.Static("/", filepath.Join(filepath.Base("."), "api", "views"))

	return app
}

func main() {
	// Load environment variables before anything else
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing with environment variables")
	}

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
