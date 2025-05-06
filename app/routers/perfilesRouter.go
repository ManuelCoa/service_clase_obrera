package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PerfilesRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	perfilesGroup := api.Group("/perfiles")

	perfilesGroup.Use(middleware.CorsMiddleware(corsOptions))
	perfilesGroup.Use(middleware.AuthRequired)
	perfilesGroup.Use(middleware.CheckPermissions)

	perfilesGroup.Get("/", controllers.GetPerfiles)
	perfilesGroup.Post("/", controllers.PostPerfiles)
}
