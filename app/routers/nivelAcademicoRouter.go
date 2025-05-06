package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func NivelAcademicoRoutes(app *fiber.App) {
	api := app.Group("/api")
	
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	nivelAcademicoGroup := api.Group("/nivelacademico")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	
	nivelAcademicoGroup.Use(middleware.CorsMiddleware(corsOptions))
	nivelAcademicoGroup.Use(middleware.AuthRequired)
	nivelAcademicoGroup.Use(middleware.CheckPermissions)
	

	nivelAcademicoGroup.Get("/", controllers.GetNivelAcademico)
	nivelAcademicoGroup.Get("/:id", controllers.GetNivelAcademicoID)
	nivelAcademicoGroup.Post("/", controllers.PostNivelAcademico)
	nivelAcademicoGroup.Put("/", controllers.PutNivelAcademico)
	nivelAcademicoGroup.Delete("/:id", controllers.DeleteNivelAcademico)
}