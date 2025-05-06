package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ConstanciaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	constanciaGroup := api.Group("/constancia-trabajo")

	constanciaGroup.Use(middleware.CorsMiddleware(corsOptions))
	constanciaGroup.Use(middleware.AuthRequired)
	constanciaGroup.Use(middleware.CheckPermissions)

	constanciaGroup.Get("/", controllers.GetConstancias)
	constanciaGroup.Get("/:id", controllers.GetConstanciaID)
	constanciaGroup.Post("/", controllers.PostConstancia)
	constanciaGroup.Put("/", controllers.PutConstancia)
	constanciaGroup.Delete("/:id", controllers.DeleteConstancia)
}