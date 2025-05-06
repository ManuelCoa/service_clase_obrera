package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func DatosAcademicosRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	datosAcademicosGroup := api.Group("/datos-academicos")

	datosAcademicosGroup.Use(middleware.CorsMiddleware(corsOptions))
	datosAcademicosGroup.Use(middleware.AuthRequired)
	datosAcademicosGroup.Use(middleware.CheckPermissions)

	datosAcademicosGroup.Get("/", controllers.GetDatosAcademicos)
	datosAcademicosGroup.Get("/:id", controllers.GetDatosAcademicosID)
	datosAcademicosGroup.Get("/por-obrero/:cedula", controllers.GetDatosAcademicosPorObrero)
	datosAcademicosGroup.Post("/", controllers.PostDatosAcademicos)
	datosAcademicosGroup.Put("/", controllers.PutDatosAcademicos)
	datosAcademicosGroup.Delete("/:id", controllers.DeleteDatosAcademicos)
}