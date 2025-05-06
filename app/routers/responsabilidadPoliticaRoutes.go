package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ResponsabilidadPoliticaRoutes(app *fiber.App) {
	api := app.Group("/api")

	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	responsabilidasPoliticaGroup := api.Group("/responsabilidad")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.

	responsabilidasPoliticaGroup.Use(middleware.CorsMiddleware(corsOptions))
	responsabilidasPoliticaGroup.Use(middleware.AuthRequired)
	responsabilidasPoliticaGroup.Use(middleware.CheckPermissions)

	responsabilidasPoliticaGroup.Get("/", controllers.GetResponsabilidadPolitica)
	responsabilidasPoliticaGroup.Get("/:id", controllers.GetResponsabilidadPoliticaID)
	responsabilidasPoliticaGroup.Post("/", controllers.PostResponsabilidadPolitica)
	responsabilidasPoliticaGroup.Put("/", controllers.PutResponsabilidadPolitica)
	responsabilidasPoliticaGroup.Delete("/:id", controllers.DeleteResponsabilidadPolitica)
}
