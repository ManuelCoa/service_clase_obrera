package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ResponsabilidadComunalRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	rcGroup := api.Group("/responsabilidad-comunal")

	rcGroup.Use(middleware.CorsMiddleware(corsOptions))
	rcGroup.Use(middleware.AuthRequired)
	rcGroup.Use(middleware.CheckPermissions)

	rcGroup.Get("/", controllers.GetResponsabilidadComunal)
	rcGroup.Get("/:id", controllers.GetResponsabilidadComunalID)
	rcGroup.Post("/", controllers.PostResponsabilidadComunal)
	rcGroup.Put("/", controllers.PutResponsabilidadComunal)
	rcGroup.Delete("/:id", controllers.DeleteResponsabilidadComunal)
}