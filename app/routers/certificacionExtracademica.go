package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func CertificacionRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	certificacionGroup := api.Group("/certificacion-extracademica")

	certificacionGroup.Use(middleware.CorsMiddleware(corsOptions))
	certificacionGroup.Use(middleware.AuthRequired)
	certificacionGroup.Use(middleware.CheckPermissions)

	certificacionGroup.Get("/", controllers.GetCertificaciones)
	certificacionGroup.Get("/:id", controllers.GetCertificacionID)
	certificacionGroup.Post("/", controllers.PostCertificacion)
	certificacionGroup.Put("/", controllers.PutCertificacion)
	certificacionGroup.Delete("/:id", controllers.DeleteCertificacion)
}