package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AsistenciaBiometricaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	asistenciaBiometricaGroup := api.Group("/asistencia-biometrica")

	asistenciaBiometricaGroup.Use(middleware.CorsMiddleware(corsOptions))
	asistenciaBiometricaGroup.Use(middleware.AuthRequired)
	asistenciaBiometricaGroup.Use(middleware.CheckPermissions)

	asistenciaBiometricaGroup.Get("/", controllers.GetAsistenciasBiometricas)
	asistenciaBiometricaGroup.Get("/:id", controllers.GetAsistenciaBiometricaID)
	asistenciaBiometricaGroup.Post("/", controllers.PostAsistenciaBiometrica)
	asistenciaBiometricaGroup.Put("/", controllers.PutAsistenciaBiometrica)
	asistenciaBiometricaGroup.Delete("/:id", controllers.DeleteAsistenciaBiometrica)
}