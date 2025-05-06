package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AsistenciaEventosRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	asistenciaGroup := api.Group("/asistencia-evento")

	// Middlewares
	asistenciaGroup.Use(middleware.CorsMiddleware(corsOptions))
	asistenciaGroup.Use(middleware.AuthRequired)
	asistenciaGroup.Use(middleware.CheckPermissions)

	// Rutas
	asistenciaGroup.Get("/", controllers.GetAsistenciaEvento)
	asistenciaGroup.Get("/:id", controllers.GetAsistenciaEventoID)
	asistenciaGroup.Post("/", controllers.PostAsistenciaEvento)
	asistenciaGroup.Put("/", controllers.PutAsistenciaEvento)
	asistenciaGroup.Delete("/:id", controllers.DeleteAsistenciaEvento)
}