package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TipoEnfermedadRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	teGroup := api.Group("/tipo-enfermedad")

	teGroup.Use(middleware.CorsMiddleware(corsOptions))
	teGroup.Use(middleware.AuthRequired)
	teGroup.Use(middleware.CheckPermissions)

	teGroup.Get("/", controllers.GetTipoEnfermedad)
	teGroup.Get("/:id", controllers.GetTipoEnfermedadID)
	teGroup.Post("/", controllers.PostTipoEnfermedad)
	teGroup.Put("/", controllers.PutTipoEnfermedad)
	teGroup.Delete("/:id", controllers.DeleteTipoEnfermedad)
}