package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TitulosAcademicosRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	tituloGroup := api.Group("/titulo-academico")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	tituloGroup.Use(middleware.CorsMiddleware(corsOptions))
	tituloGroup.Use(middleware.AuthRequired)
	tituloGroup.Use(middleware.CheckPermissions)

	tituloGroup.Get("/", controllers.GetTituloAcademico)
	tituloGroup.Get("/:id", controllers.GetTituloAcademicoID)
	tituloGroup.Post("/", controllers.PostTituloAcademico)
	tituloGroup.Put("/", controllers.PutTituloAcademico)
	tituloGroup.Delete("/:id", controllers.DeleteTituloAcademico)
}