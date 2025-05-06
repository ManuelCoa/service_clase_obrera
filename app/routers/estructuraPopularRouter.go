package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func EstructuraPopularRoutes(app *fiber.App) {
	api := app.Group("/api")
	
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	
	EstructuraPopularGroup := api.Group("/estructurapopular")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	
	EstructuraPopularGroup.Use(middleware.CorsMiddleware(corsOptions))
	EstructuraPopularGroup.Use(middleware.AuthRequired)
	EstructuraPopularGroup.Use(middleware.CheckPermissions)
	

	EstructuraPopularGroup.Get("/", controllers.GetEstructuraPopular)
	EstructuraPopularGroup.Get("/:id", controllers.GetEstructuraPopularID)
	EstructuraPopularGroup.Post("/", controllers.PostEstructuraPopular)
	EstructuraPopularGroup.Put("/", controllers.PutEstructuraPopular)
	EstructuraPopularGroup.Delete("/:id", controllers.DeleteEstructuraPopular)
}