package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func EventosPopularesRoutes(app *fiber.App) {
	api := app.Group("/api")
	
	configCORS := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	eventoGroup := api.Group("/evento-popular")
	eventoGroup.Use(middleware.CorsMiddleware(configCORS))
	eventoGroup.Use(middleware.AuthRequired)
	eventoGroup.Use(middleware.CheckPermissions)


	// Rutas
	eventoGroup.Get("/", controllers.GetEventoPopular)
	eventoGroup.Get("/:id", controllers.GetEventoPopularID)
	eventoGroup.Post("/", controllers.PostEventoPopular)
	eventoGroup.Put("/", controllers.PutEventoPopular)
	eventoGroup.Delete("/:id", controllers.DeleteEventoPopular)
}