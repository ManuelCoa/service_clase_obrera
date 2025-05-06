package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SessionRoutes(app *fiber.App) {

	api := app.Group("/api")
	
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	
	sessionGroup := api.Group("/session")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	sessionGroup.Use(middleware.CorsMiddleware(corsOptions))
	
	sessionGroup.Post("/login", controllers.PostLogin)
	sessionGroup.Post("/logout", controllers.PostLogout)
}
