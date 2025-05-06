package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ParentescoRoutes(app *fiber.App) {
	api := app.Group("/api")

		corsOptions := middleware.CorsOptions{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
			AllowedHeaders:   []string{"Content-Type", "Authorization"},
			AllowCredentials: true,
			ExposedHeaders:   []string{"X-Custom-Header"},
			MaxAge:           3600,
		}

	ParentescoGroup := api.Group("/parentesco")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.

		ParentescoGroup.Use(middleware.CorsMiddleware(corsOptions))
		ParentescoGroup.Use(middleware.AuthRequired)
		ParentescoGroup.Use(middleware.CheckPermissions)


	ParentescoGroup.Get("/", controllers.GetParentesco)
	ParentescoGroup.Get("/:id", controllers.GetParentescoID)
	ParentescoGroup.Post("/", controllers.PostParentesco)
	ParentescoGroup.Put("/", controllers.PutParentesco)
	ParentescoGroup.Delete("/:id", controllers.DeleteParentesco)
}
