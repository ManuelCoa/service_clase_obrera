package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ConsejoComunalRoutes(app *fiber.App) {
	api := app.Group("/api")
	
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	
	consejoComunalGroup := api.Group("/consejocomunal")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	
	consejoComunalGroup.Use(middleware.CorsMiddleware(corsOptions))
	consejoComunalGroup.Use(middleware.AuthRequired)
	consejoComunalGroup.Use(middleware.CheckPermissions)
	

	consejoComunalGroup.Get("/", controllers.GetConsejoComunal)
	consejoComunalGroup.Get("/:id", controllers.GetConsejoComunalID)
	consejoComunalGroup.Post("/", controllers.PostConsejoComunal)
	consejoComunalGroup.Put("/", controllers.PutConsejoComunal)
	consejoComunalGroup.Delete("/:id", controllers.DeleteConsejoComunal)
}