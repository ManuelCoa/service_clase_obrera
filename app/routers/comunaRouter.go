package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ComunaRoutes(app *fiber.App) {
	api := app.Group("/api")
	
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	
	comunaGroup := api.Group("/comuna")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	
	comunaGroup.Use(middleware.CorsMiddleware(corsOptions))
	comunaGroup.Use(middleware.AuthRequired)
	comunaGroup.Use(middleware.CheckPermissions)
	

	comunaGroup.Get("/", controllers.GetComuna)
	comunaGroup.Get("/:id", controllers.GetComunaID)
	comunaGroup.Post("/", controllers.PostComuna)
	comunaGroup.Put("/", controllers.PutComuna)
	comunaGroup.Delete("/:id", controllers.DeleteComuna)
}