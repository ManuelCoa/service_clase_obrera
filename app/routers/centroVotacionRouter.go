package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func CentrosVotacionRoutes(app *fiber.App) {
	api := app.Group("/api")
	
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	
	centroVotacionGroup := api.Group("/centrovotacion")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	
	centroVotacionGroup.Use(middleware.CorsMiddleware(corsOptions))
	centroVotacionGroup.Use(middleware.AuthRequired)
	centroVotacionGroup.Use(middleware.CheckPermissions)
	

	centroVotacionGroup.Get("/", controllers.GetCentroVotacion)
	centroVotacionGroup.Get("/:id", controllers.GetCentroVotacionID)
	centroVotacionGroup.Post("/", controllers.PostCentroVotacion)
	centroVotacionGroup.Put("/", controllers.PutCentroVotacion)
	centroVotacionGroup.Delete("/:id", controllers.DeleteCentroVotacion)
}
