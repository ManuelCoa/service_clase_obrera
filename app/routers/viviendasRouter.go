package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ViviendaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	viviendaGroup := api.Group("/vivienda")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	viviendaGroup.Use(middleware.CorsMiddleware(corsOptions))
	viviendaGroup.Use(middleware.AuthRequired)
	viviendaGroup.Use(middleware.CheckPermissions)

	viviendaGroup.Get("/", controllers.GetVivienda)
	viviendaGroup.Get("/:id", controllers.GetViviendaID)
	viviendaGroup.Post("/", controllers.PostVivienda)
	viviendaGroup.Put("/", controllers.PutVivienda)
	viviendaGroup.Delete("/:id", controllers.DeleteVivienda)
}