package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TipoHabilidadRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	thGroup := api.Group("/tipo-habilidad")

	thGroup.Use(middleware.CorsMiddleware(corsOptions))
	thGroup.Use(middleware.AuthRequired)
	thGroup.Use(middleware.CheckPermissions)

	thGroup.Get("/", controllers.GetTipoHabilidad)
	thGroup.Get("/:id", controllers.GetTipoHabilidadID)
	thGroup.Post("/", controllers.PostTipoHabilidad)
	thGroup.Put("/", controllers.PutTipoHabilidad)
	thGroup.Delete("/:id", controllers.DeleteTipoHabilidad)
}