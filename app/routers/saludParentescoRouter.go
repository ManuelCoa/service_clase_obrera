package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SaludParentescoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	saludParentescoGroup := api.Group("/salud-parentesco")

	saludParentescoGroup.Use(middleware.CorsMiddleware(corsOptions))
	saludParentescoGroup.Use(middleware.AuthRequired)
	saludParentescoGroup.Use(middleware.CheckPermissions)

	saludParentescoGroup.Get("/", controllers.GetSaludParentescos)
	saludParentescoGroup.Get("/:id", controllers.GetSaludParentescoID)
	saludParentescoGroup.Post("/", controllers.PostSaludParentesco)
	saludParentescoGroup.Put("/", controllers.PutSaludParentesco)
	saludParentescoGroup.Delete("/:id", controllers.DeleteSaludParentesco)
}