package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func HorarioRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	horarioGroup := api.Group("/horario")

	horarioGroup.Use(middleware.CorsMiddleware(corsOptions))
	horarioGroup.Use(middleware.AuthRequired)
	horarioGroup.Use(middleware.CheckPermissions)

	horarioGroup.Get("/", controllers.GetHorario)
	horarioGroup.Get("/:id", controllers.GetHorarioID)
	horarioGroup.Post("/", controllers.PostHorario)
	horarioGroup.Put("/", controllers.PutHorario)
	horarioGroup.Delete("/:id", controllers.DeleteHorario)
}