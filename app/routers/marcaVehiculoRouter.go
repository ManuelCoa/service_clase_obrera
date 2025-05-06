package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func MarcaVehiculoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	mvGroup := api.Group("/marca-vehiculo")

	mvGroup.Use(middleware.CorsMiddleware(corsOptions))
	mvGroup.Use(middleware.AuthRequired)
	mvGroup.Use(middleware.CheckPermissions)

	mvGroup.Get("/", controllers.GetMarcaVehiculo)
	mvGroup.Get("/:id", controllers.GetMarcaVehiculoID)
	mvGroup.Post("/", controllers.PostMarcaVehiculo)
	mvGroup.Put("/", controllers.PutMarcaVehiculo)
	mvGroup.Delete("/:id", controllers.DeleteMarcaVehiculo)
}