package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func DireccionParentescoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	direccionParentescoGroup := api.Group("/direccion-parentesco")

	direccionParentescoGroup.Use(middleware.CorsMiddleware(corsOptions))
	direccionParentescoGroup.Use(middleware.AuthRequired)
	direccionParentescoGroup.Use(middleware.CheckPermissions)

	direccionParentescoGroup.Get("/", controllers.GetDireccionParentesco)
	direccionParentescoGroup.Get("/:id", controllers.GetDireccionParentescoID)
	direccionParentescoGroup.Post("/", controllers.PostDireccionParentesco)
	direccionParentescoGroup.Put("/", controllers.PutDireccionParentesco)
	direccionParentescoGroup.Delete("/:id", controllers.DeleteDireccionParentesco)
}