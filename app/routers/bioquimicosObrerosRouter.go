package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func BioquimicosObreroRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	bioquimicosObreroGroup := api.Group("/bioquimicos-obrero")

	bioquimicosObreroGroup.Use(middleware.CorsMiddleware(corsOptions))
	bioquimicosObreroGroup.Use(middleware.AuthRequired)
	bioquimicosObreroGroup.Use(middleware.CheckPermissions)

	bioquimicosObreroGroup.Get("/", controllers.GetBioquimicosObreros)
	bioquimicosObreroGroup.Get("/:id", controllers.GetBioquimicosObreroID)
	bioquimicosObreroGroup.Post("/", controllers.PostBioquimicosObrero)
	bioquimicosObreroGroup.Put("/", controllers.PutBioquimicosObrero)
	bioquimicosObreroGroup.Delete("/:id", controllers.DeleteBioquimicosObrero)
}