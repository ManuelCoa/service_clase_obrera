package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func CargoPublicoRoutes(app *fiber.App) {
	api := app.Group("/api")

	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	cargoGroup := api.Group("/cargopublico")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	cargoGroup.Use(middleware.CorsMiddleware(corsOptions))
	cargoGroup.Use(middleware.AuthRequired)
	cargoGroup.Use(middleware.CheckPermissions)

	cargoGroup.Get("/", controllers.GetCargoPublico)
	cargoGroup.Get("/:id", controllers.GetCargoPublicoID)
	cargoGroup.Post("/", controllers.PostCargoPublico)
	cargoGroup.Put("/", controllers.PutCargoPublico)
	cargoGroup.Delete("/:id", controllers.DeleteCargoPublico)
}
