package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PoliticoObreroRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	politicoObreroGroup := api.Group("/politico-obrero")

	politicoObreroGroup.Use(middleware.CorsMiddleware(corsOptions))
	politicoObreroGroup.Use(middleware.AuthRequired)
	politicoObreroGroup.Use(middleware.CheckPermissions)

	politicoObreroGroup.Get("/", controllers.GetPoliticosObreros)
	politicoObreroGroup.Get("/:id", controllers.GetPoliticoObreroID)
	politicoObreroGroup.Post("/", controllers.PostPoliticoObrero)
	politicoObreroGroup.Put("/", controllers.PutPoliticoObrero)
	politicoObreroGroup.Delete("/:id", controllers.DeletePoliticoObrero)
}