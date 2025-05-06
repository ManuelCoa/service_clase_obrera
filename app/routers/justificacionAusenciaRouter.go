package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func JustificacionAusenciaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	justificacionGroup := api.Group("/justificacion-ausencia")

	justificacionGroup.Use(middleware.CorsMiddleware(corsOptions))
	justificacionGroup.Use(middleware.AuthRequired)
	justificacionGroup.Use(middleware.CheckPermissions)

	justificacionGroup.Get("/", controllers.GetJustificacionesAusencia)
	justificacionGroup.Get("/:id", controllers.GetJustificacionAusenciaID)
	justificacionGroup.Post("/", controllers.PostJustificacionAusencia)
	justificacionGroup.Put("/", controllers.PutJustificacionAusencia)
	justificacionGroup.Delete("/:id", controllers.DeleteJustificacionAusencia)
}