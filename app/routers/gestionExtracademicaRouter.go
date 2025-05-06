package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func GestionExtracademicaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	gestionExtracademicaGroup := api.Group("/gestion-extracademica")

	gestionExtracademicaGroup.Use(middleware.CorsMiddleware(corsOptions))
	gestionExtracademicaGroup.Use(middleware.AuthRequired)
	gestionExtracademicaGroup.Use(middleware.CheckPermissions)

	gestionExtracademicaGroup.Get("/", controllers.GetGestionesExtracademicas)
	gestionExtracademicaGroup.Get("/:id", controllers.GetGestionExtracademicaID)
	gestionExtracademicaGroup.Post("/", controllers.PostGestionExtracademica)
	gestionExtracademicaGroup.Put("/", controllers.PutGestionExtracademica)
	gestionExtracademicaGroup.Delete("/:id", controllers.DeleteGestionExtracademica)
}