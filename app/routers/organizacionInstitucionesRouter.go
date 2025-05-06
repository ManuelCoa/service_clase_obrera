package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func InstitucionesRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	institucionesGroup := api.Group("/instituciones")
	institucionesGroup.Use(middleware.CorsMiddleware(corsOptions))
	institucionesGroup.Use(middleware.AuthRequired)
	institucionesGroup.Use(middleware.CheckPermissions)

	institucionesGroup.Get("/", controllers.GetInstituciones)
	institucionesGroup.Get("/:id", controllers.GetInstitucionesID)
	institucionesGroup.Post("/", controllers.PostInstituciones)
	institucionesGroup.Put("/", controllers.PutInstituciones)
	institucionesGroup.Delete("/:id", controllers.DeleteInstituciones)
}
