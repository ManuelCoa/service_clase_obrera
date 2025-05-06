package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TipoSueloViviendaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	tsvGroup := api.Group("/tipo-suelo-vivienda")

	tsvGroup.Use(middleware.CorsMiddleware(corsOptions))
	tsvGroup.Use(middleware.AuthRequired)
	tsvGroup.Use(middleware.CheckPermissions)

	tsvGroup.Get("/", controllers.GetTipoSueloVivienda)
	tsvGroup.Get("/:id", controllers.GetTipoSueloViviendaID)
	tsvGroup.Post("/", controllers.PostTipoSueloVivienda)
	tsvGroup.Put("/", controllers.PutTipoSueloVivienda)
	tsvGroup.Delete("/:id", controllers.DeleteTipoSueloVivienda)
}