package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func GestionParentescoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	parentescoGroup := api.Group("/gestion-parentesco")

	parentescoGroup.Use(middleware.CorsMiddleware(corsOptions))
	parentescoGroup.Use(middleware.AuthRequired)
	parentescoGroup.Use(middleware.CheckPermissions)

	parentescoGroup.Get("/", controllers.GetGestionParentescos)
	parentescoGroup.Get("/:id", controllers.GetGestionParentescoID)
	parentescoGroup.Post("/", controllers.PostGestionParentesco)
	parentescoGroup.Put("/", controllers.PutGestionParentesco)
	parentescoGroup.Delete("/:id", controllers.DeleteGestionParentesco)
}