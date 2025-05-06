package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SolicitudIngresoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	solicitudIngresoGroup := api.Group("/solicitud-ingreso")

	solicitudIngresoGroup.Use(middleware.CorsMiddleware(corsOptions))
	solicitudIngresoGroup.Use(middleware.AuthRequired)
	solicitudIngresoGroup.Use(middleware.CheckPermissions)

	solicitudIngresoGroup.Get("/", controllers.GetSolicitudesIngreso)
	solicitudIngresoGroup.Get("/:id", controllers.GetSolicitudIngresoID)
	solicitudIngresoGroup.Post("/", controllers.PostSolicitudIngreso)
	solicitudIngresoGroup.Put("/", controllers.PutSolicitudIngreso)
	solicitudIngresoGroup.Delete("/:id", controllers.DeleteSolicitudIngreso)
}