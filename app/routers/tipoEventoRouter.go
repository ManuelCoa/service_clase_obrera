package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TipoEventosRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	tipoEventoGroup := api.Group("/tipo-evento")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	tipoEventoGroup.Use(middleware.CorsMiddleware(corsOptions))
	tipoEventoGroup.Use(middleware.AuthRequired)
	tipoEventoGroup.Use(middleware.CheckPermissions)

	tipoEventoGroup.Get("/", controllers.GetTipoEvento)
	tipoEventoGroup.Get("/:id", controllers.GetTipoEventoID)
	tipoEventoGroup.Post("/", controllers.PostTipoEvento)
	tipoEventoGroup.Put("/", controllers.PutTipoEvento)
	tipoEventoGroup.Delete("/:id", controllers.DeleteTipoEvento)
}