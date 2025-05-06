package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TipoViviendaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	tvGroup := api.Group("/tipo-vivienda")

	tvGroup.Use(middleware.CorsMiddleware(corsOptions))
	tvGroup.Use(middleware.AuthRequired)
	tvGroup.Use(middleware.CheckPermissions)

	tvGroup.Get("/", controllers.GetTipoVivienda)
	tvGroup.Get("/:id", controllers.GetTipoViviendaID)
	tvGroup.Post("/", controllers.PostTipoVivienda)
	tvGroup.Put("/", controllers.PutTipoVivienda)
	tvGroup.Delete("/:id", controllers.DeleteTipoVivienda)
}