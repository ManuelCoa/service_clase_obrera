package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func DatosLaboralesRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	datosLaboralesGroup := api.Group("/datos-laborales")

	datosLaboralesGroup.Use(middleware.CorsMiddleware(corsOptions))
	datosLaboralesGroup.Use(middleware.AuthRequired)
	datosLaboralesGroup.Use(middleware.CheckPermissions)

	datosLaboralesGroup.Get("/", controllers.GetDatosLaborales)
	datosLaboralesGroup.Get("/:id", controllers.GetDatosLaboralesID)
	datosLaboralesGroup.Get("/por-obrero/:cedula", controllers.GetDatosLaboralesPorObrero)
	datosLaboralesGroup.Post("/", controllers.PostDatosLaborales)
	datosLaboralesGroup.Put("/", controllers.PutDatosLaborales)
	datosLaboralesGroup.Delete("/:id", controllers.DeleteDatosLaborales)
}