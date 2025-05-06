package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TipoFormacionExtracademicaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	tfeGroup := api.Group("/tipo-formacion-extracademica")

	tfeGroup.Use(middleware.CorsMiddleware(corsOptions))
	tfeGroup.Use(middleware.AuthRequired)
	tfeGroup.Use(middleware.CheckPermissions)

	tfeGroup.Get("/", controllers.GetTipoFormacionExtracademica)
	tfeGroup.Get("/:id", controllers.GetTipoFormacionExtracademicaID)
	tfeGroup.Post("/", controllers.PostTipoFormacionExtracademica)
	tfeGroup.Put("/", controllers.PutTipoFormacionExtracademica)
	tfeGroup.Delete("/:id", controllers.DeleteTipoFormacionExtracademica)
}