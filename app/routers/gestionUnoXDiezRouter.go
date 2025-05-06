package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func GestionUnoxDiezRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	gestionGroup := api.Group("/gestion-1x10")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	gestionGroup.Use(middleware.CorsMiddleware(corsOptions))
	gestionGroup.Use(middleware.AuthRequired)
	gestionGroup.Use(middleware.CheckPermissions)

	gestionGroup.Get("/", controllers.GetGestionUnoxDiez)
	gestionGroup.Get("/:id", controllers.GetGestionUnoxDiezID)
	gestionGroup.Post("/", controllers.PostGestionUnoxDiez)
	gestionGroup.Put("/", controllers.PutGestionUnoxDiez)
	gestionGroup.Delete("/:id", controllers.DeleteGestionUnoxDiez)
}