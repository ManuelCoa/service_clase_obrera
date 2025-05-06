package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AsistenciaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	asistenciaGroup := api.Group("/asistencia")

	// Se aplica el middleware de CORS
	asistenciaGroup.Use(middleware.CorsMiddleware(corsOptions))
	asistenciaGroup.Use(middleware.AuthRequired)
	asistenciaGroup.Use(middleware.CheckPermissions)

	asistenciaGroup.Get("/", controllers.GetAsistencias)
	asistenciaGroup.Get("/:id", controllers.GetAsistenciaID)
	asistenciaGroup.Post("/", controllers.PostAsistencia)
	asistenciaGroup.Put("/", controllers.PutAsistencia)
	asistenciaGroup.Delete("/:id", controllers.DeleteAsistencia)
}