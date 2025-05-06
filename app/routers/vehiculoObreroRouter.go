package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func VehiculoObreroRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	vehiculoGroup := api.Group("/vehiculo-obrero")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	vehiculoGroup.Use(middleware.CorsMiddleware(corsOptions))
	vehiculoGroup.Use(middleware.AuthRequired)
	vehiculoGroup.Use(middleware.CheckPermissions)

	vehiculoGroup.Get("/", controllers.GetVehiculoObrero)
	vehiculoGroup.Get("/:id", controllers.GetVehiculoObreroID)
	vehiculoGroup.Post("/", controllers.PostVehiculoObrero)
	vehiculoGroup.Put("/", controllers.PutVehiculoObrero)
	vehiculoGroup.Delete("/:id", controllers.DeleteVehiculoObrero)
}