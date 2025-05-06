package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AmonestacionesRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	amonestacionesGroup := api.Group("/amonestaciones")

	// Se aplica el middleware de CORS
	amonestacionesGroup.Use(middleware.CorsMiddleware(corsOptions))
	amonestacionesGroup.Use(middleware.AuthRequired)
	amonestacionesGroup.Use(middleware.CheckPermissions)

	amonestacionesGroup.Get("/", controllers.GetAmonestaciones)
	amonestacionesGroup.Get("/:id", controllers.GetAmonestacionID)
	amonestacionesGroup.Post("/", controllers.PostAmonestacion)
	amonestacionesGroup.Put("/", controllers.PutAmonestacion)
	amonestacionesGroup.Delete("/:id", controllers.DeleteAmonestacion)
}