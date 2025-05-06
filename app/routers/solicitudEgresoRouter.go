package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SolicitudEgresoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	solicitudEgresoGroup := api.Group("/solicitud-egreso")

	solicitudEgresoGroup.Use(middleware.CorsMiddleware(corsOptions))
	solicitudEgresoGroup.Use(middleware.AuthRequired)
	solicitudEgresoGroup.Use(middleware.CheckPermissions)

	solicitudEgresoGroup.Get("/", controllers.GetSolicitudesEgreso)
	solicitudEgresoGroup.Get("/:id", controllers.GetSolicitudEgresoID)
	solicitudEgresoGroup.Post("/", controllers.PostSolicitudEgreso)
	solicitudEgresoGroup.Put("/", controllers.PutSolicitudEgreso)
	solicitudEgresoGroup.Delete("/:id", controllers.DeleteSolicitudEgreso)
}