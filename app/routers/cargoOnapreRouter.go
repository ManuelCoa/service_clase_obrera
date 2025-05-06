package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func CargoOnapreRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	cargoOnapreGroup := api.Group("/cargoonapre")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	cargoOnapreGroup.Use(middleware.CorsMiddleware(corsOptions))
	cargoOnapreGroup.Use(middleware.AuthRequired)
	cargoOnapreGroup.Use(middleware.CheckPermissions)

	cargoOnapreGroup.Get("/", controllers.GetCargoOnapre)
	cargoOnapreGroup.Get("/:id", controllers.GetCargoOnapreID)
	cargoOnapreGroup.Post("/", controllers.PostCargoOnapre)
	cargoOnapreGroup.Put("/", controllers.PutCargoOnapre)
	cargoOnapreGroup.Delete("/:id", controllers.DeleteCargoOnapre)
}
