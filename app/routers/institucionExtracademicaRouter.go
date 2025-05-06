package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func InstitucionExtracademicaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	ieGroup := api.Group("/institucion-extracademica")

	ieGroup.Use(middleware.CorsMiddleware(corsOptions))
	ieGroup.Use(middleware.AuthRequired)
	ieGroup.Use(middleware.CheckPermissions)

	ieGroup.Get("/", controllers.GetInstitucionExtracademica)
	ieGroup.Get("/:id", controllers.GetInstitucionExtracademicaID)
	ieGroup.Post("/", controllers.PostInstitucionExtracademica)
	ieGroup.Put("/", controllers.PutInstitucionExtracademica)
	ieGroup.Delete("/:id", controllers.DeleteInstitucionExtracademica)
}