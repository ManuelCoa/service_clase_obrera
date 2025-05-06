package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func MovimientosSocialesRoutes(app *fiber.App) {
	api := app.Group("/api")
	
		corsOptions := middleware.CorsOptions{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
			AllowedHeaders:   []string{"Content-Type", "Authorization"},
			AllowCredentials: true,
			ExposedHeaders:   []string{"X-Custom-Header"},
			MaxAge:           3600,
		}
	
	movimientosScialesGroup := api.Group("/movimientosociales")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	
	movimientosScialesGroup.Use(middleware.CorsMiddleware(corsOptions))
	movimientosScialesGroup.Use(middleware.AuthRequired)
	movimientosScialesGroup.Use(middleware.CheckPermissions)
	

	movimientosScialesGroup.Get("/", controllers.GetMovimientosSociales)
	movimientosScialesGroup.Get("/:id", controllers.GetMovimientosSocialesID)
	movimientosScialesGroup.Post("/", controllers.PostMovimientosSociales)
	movimientosScialesGroup.Put("/", controllers.PutMovimientosSociales)
	movimientosScialesGroup.Delete("/:id", controllers.DeleteMovimientosSociales)
}
