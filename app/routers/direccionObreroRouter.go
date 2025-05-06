package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func DireccionObreroRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	direccionObreroGroup := api.Group("/direccion-obrero")

	direccionObreroGroup.Use(middleware.CorsMiddleware(corsOptions))
	direccionObreroGroup.Use(middleware.AuthRequired)
	direccionObreroGroup.Use(middleware.CheckPermissions)

	direccionObreroGroup.Get("/", controllers.GetDireccionesObrero)
	direccionObreroGroup.Get("/:id", controllers.GetDireccionObreroID)
	direccionObreroGroup.Post("/", controllers.PostDireccionObrero)
	direccionObreroGroup.Put("/", controllers.PutDireccionObrero)
	direccionObreroGroup.Delete("/:id", controllers.DeleteDireccionObrero)
}