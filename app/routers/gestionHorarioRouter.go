package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func GestionHorarioRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	gestionHorarioGroup := api.Group("/gestion-horario")

	gestionHorarioGroup.Use(middleware.CorsMiddleware(corsOptions))
	gestionHorarioGroup.Use(middleware.AuthRequired)
	gestionHorarioGroup.Use(middleware.CheckPermissions)

	gestionHorarioGroup.Get("/", controllers.GetGestionHorarios)
	gestionHorarioGroup.Get("/:id", controllers.GetGestionHorarioID)
	gestionHorarioGroup.Post("/", controllers.PostGestionHorario)
	gestionHorarioGroup.Put("/", controllers.PutGestionHorario)
	gestionHorarioGroup.Delete("/:id", controllers.DeleteGestionHorario)
}