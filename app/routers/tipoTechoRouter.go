package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TipoTechoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	ttGroup := api.Group("/tipo-techo")

	ttGroup.Use(middleware.CorsMiddleware(corsOptions))
	ttGroup.Use(middleware.AuthRequired)
	ttGroup.Use(middleware.CheckPermissions)

	ttGroup.Get("/", controllers.GetTipoTecho)
	ttGroup.Get("/:id", controllers.GetTipoTechoID)
	ttGroup.Post("/", controllers.PostTipoTecho)
	ttGroup.Put("/", controllers.PutTipoTecho)
	ttGroup.Delete("/:id", controllers.DeleteTipoTecho)
}