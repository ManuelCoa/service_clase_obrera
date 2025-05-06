package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AntecedentesMedicosRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	antecedentesMedicosGroup := api.Group("/antecedentes-medicos")

	// Se aplica el middleware de CORS
	antecedentesMedicosGroup.Use(middleware.CorsMiddleware(corsOptions))
	antecedentesMedicosGroup.Use(middleware.AuthRequired)
	antecedentesMedicosGroup.Use(middleware.CheckPermissions)

	antecedentesMedicosGroup.Get("/", controllers.GetAntecedentesMedicos)
	antecedentesMedicosGroup.Get("/:id", controllers.GetAntecedenteMedicoID)
	antecedentesMedicosGroup.Post("/", controllers.PostAntecedenteMedico)
	antecedentesMedicosGroup.Put("/", controllers.PutAntecedenteMedico)
	antecedentesMedicosGroup.Delete("/:id", controllers.DeleteAntecedenteMedico)
}