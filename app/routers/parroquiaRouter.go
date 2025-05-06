package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ParroquiasRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	parroquiaGroup := api.Group("/parroquia") 
	
	parroquiaGroup.Use(middleware.CorsMiddleware(corsOptions))
	parroquiaGroup.Use(middleware.AuthRequired)
	parroquiaGroup.Use(middleware.CheckPermissions)

	parroquiaGroup.Get("/", controllers.GetParroquia)
	parroquiaGroup.Get("/:id", controllers.GetParroquiaID)
	parroquiaGroup.Post("/", controllers.PostParroquia)
	parroquiaGroup.Put("/", controllers.PutParroquia)
	parroquiaGroup.Delete("/:id", controllers.DeleteParroquia)
}
