package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func EstadoCivilRoutes(app *fiber.App) {
	api := app.Group("/api")
	
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	
	estadoCivilGroup := api.Group("/estadocivil")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	
	estadoCivilGroup.Use(middleware.CorsMiddleware(corsOptions))
	estadoCivilGroup.Use(middleware.AuthRequired)
	estadoCivilGroup.Use(middleware.CheckPermissions)
	

	estadoCivilGroup.Get("/", controllers.GetEstadoCivil)
	estadoCivilGroup.Get("/:id", controllers.GetEstadoCivilID)
	estadoCivilGroup.Post("/", controllers.PostEstadoCivil)
	estadoCivilGroup.Put("/", controllers.PutEstadoCivil)
	estadoCivilGroup.Delete("/:id", controllers.DeleteEstadoCivil)
}