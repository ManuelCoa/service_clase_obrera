package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegistroTalentoHumanoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	talentoHumanoGroup := api.Group("/talento-humano")

	talentoHumanoGroup.Use(middleware.CorsMiddleware(corsOptions))
	talentoHumanoGroup.Use(middleware.AuthRequired)
	talentoHumanoGroup.Use(middleware.CheckPermissions)

	talentoHumanoGroup.Get("/", controllers.GetRegistrosTalentoHumano)
	talentoHumanoGroup.Get("/:id", controllers.GetRegistroTalentoHumanoID)
	talentoHumanoGroup.Post("/", controllers.PostRegistroTalentoHumano)
	talentoHumanoGroup.Put("/", controllers.PutRegistroTalentoHumano)
	talentoHumanoGroup.Delete("/:id", controllers.DeleteRegistroTalentoHumano)
}