package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func GeneroRoutes(app *fiber.App) {
	api := app.Group("/api")
	
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	
	generoGroup := api.Group("/genero")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	
	generoGroup.Use(middleware.CorsMiddleware(corsOptions))
	generoGroup.Use(middleware.AuthRequired)
	generoGroup.Use(middleware.CheckPermissions)
	

	generoGroup.Get("/", controllers.GetGenero)
	generoGroup.Get("/:id", controllers.GetGeneroID)
	generoGroup.Post("/", controllers.PostGenero)
	generoGroup.Put("/", controllers.PutGenero)
	generoGroup.Delete("/:id", controllers.DeleteGenero)
}