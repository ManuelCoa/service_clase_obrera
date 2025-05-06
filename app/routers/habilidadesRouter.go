package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func HabilidadesRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	habilidadesGroup := api.Group("/habilidades")

	habilidadesGroup.Use(middleware.CorsMiddleware(corsOptions))
	habilidadesGroup.Use(middleware.AuthRequired)
	habilidadesGroup.Use(middleware.CheckPermissions)

	habilidadesGroup.Get("/", controllers.GetHabilidades)
	habilidadesGroup.Get("/:id", controllers.GetHabilidadID)
	habilidadesGroup.Post("/", controllers.PostHabilidad)
	habilidadesGroup.Put("/", controllers.PutHabilidad)
	habilidadesGroup.Delete("/:id", controllers.DeleteHabilidad)
}