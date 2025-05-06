package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func EstadosRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	estadoGroup := api.Group("/estado")

	estadoGroup.Use(middleware.CorsMiddleware(corsOptions))
	estadoGroup.Use(middleware.AuthRequired)
	estadoGroup.Use(middleware.CheckPermissions)

	estadoGroup.Get("/", controllers.GetEstado)
	estadoGroup.Get("/:id", controllers.GetEstadoID)
	estadoGroup.Post("/", controllers.PostEstado)
	estadoGroup.Put("/", controllers.PutEstado)
	estadoGroup.Delete("/:id", controllers.DeleteEstado)
}
