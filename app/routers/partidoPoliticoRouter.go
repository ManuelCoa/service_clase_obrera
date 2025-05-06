package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PartidoPoliticoRoutes(app *fiber.App) {
	api := app.Group("/api")
	
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	
	partidoPoliticoGroup := api.Group("/partidopolitico")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	
	partidoPoliticoGroup.Use(middleware.CorsMiddleware(corsOptions))
	partidoPoliticoGroup.Use(middleware.AuthRequired)
	partidoPoliticoGroup.Use(middleware.CheckPermissions)
	

	partidoPoliticoGroup.Get("/", controllers.GetPartidoPolitico)
	partidoPoliticoGroup.Get("/:id", controllers.GetPartidoPoliticoID)
	partidoPoliticoGroup.Post("/", controllers.PostPartidoPolitico)
	partidoPoliticoGroup.Put("/", controllers.PutPartidoPolitico)
	partidoPoliticoGroup.Delete("/:id", controllers.DeletePartidoPolitico)
}