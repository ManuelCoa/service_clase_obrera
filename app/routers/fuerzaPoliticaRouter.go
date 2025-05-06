package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func FuerzaPoliticaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	fuerzaPoliticaGroup := api.Group("/fuerza-politica")

	fuerzaPoliticaGroup.Use(middleware.CorsMiddleware(corsOptions))
	fuerzaPoliticaGroup.Use(middleware.AuthRequired)
	fuerzaPoliticaGroup.Use(middleware.CheckPermissions)

	fuerzaPoliticaGroup.Get("/", controllers.GetFuerzaPolitica)
	fuerzaPoliticaGroup.Get("/:id", controllers.GetFuerzaPoliticaID)
	fuerzaPoliticaGroup.Post("/", controllers.PostFuerzaPolitica)
	fuerzaPoliticaGroup.Put("/", controllers.PutFuerzaPolitica)
	fuerzaPoliticaGroup.Delete("/:id", controllers.DeleteFuerzaPolitica)
}