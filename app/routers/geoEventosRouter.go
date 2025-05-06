package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func GeoEventosRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	geoEventoGroup := api.Group("/geo-evento")

	// Se aplica el middleware de CORS
	geoEventoGroup.Use(middleware.CorsMiddleware(corsOptions))
	geoEventoGroup.Use(middleware.AuthRequired)
	geoEventoGroup.Use(middleware.CheckPermissions)

	geoEventoGroup.Get("/", controllers.GetGeoEvento)
	geoEventoGroup.Get("/:id", controllers.GetGeoEventoID)
	geoEventoGroup.Post("/", controllers.PostGeoEvento)
	geoEventoGroup.Put("/", controllers.PutGeoEvento)
	geoEventoGroup.Delete("/:id", controllers.DeleteGeoEvento)
}