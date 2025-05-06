package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TipoParedesRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	tipoParedGroup := api.Group("/tipo-pared")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	tipoParedGroup.Use(middleware.CorsMiddleware(corsOptions))
	tipoParedGroup.Use(middleware.AuthRequired)
	tipoParedGroup.Use(middleware.CheckPermissions)

	tipoParedGroup.Get("/", controllers.GetTipoPared)
	tipoParedGroup.Get("/:id", controllers.GetTipoParedID)
	tipoParedGroup.Post("/", controllers.PostTipoPared)
	tipoParedGroup.Put("/", controllers.PutTipoPared)
	tipoParedGroup.Delete("/:id", controllers.DeleteTipoPared)
}