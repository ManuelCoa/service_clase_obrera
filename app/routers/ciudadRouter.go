package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func CiudadesRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	ciudadGroup := api.Group("/ciudad")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	ciudadGroup.Use(middleware.CorsMiddleware(corsOptions))
	ciudadGroup.Use(middleware.AuthRequired)
	ciudadGroup.Use(middleware.CheckPermissions)

	ciudadGroup.Get("/", controllers.GetCiudad)
	ciudadGroup.Get("/:id", controllers.GetCiudadID)
	ciudadGroup.Post("/", controllers.PostCiudad)
	ciudadGroup.Put("/", controllers.PutCiudad)
	ciudadGroup.Delete("/:id", controllers.DeleteCiudad)
}
