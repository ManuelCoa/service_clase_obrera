package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TipoParentescoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	tpGroup := api.Group("/tipo-parentesco")

	tpGroup.Use(middleware.CorsMiddleware(corsOptions))
	tpGroup.Use(middleware.AuthRequired)
	tpGroup.Use(middleware.CheckPermissions)

	tpGroup.Get("/", controllers.GetTipoParentesco)
	tpGroup.Get("/:id", controllers.GetTipoParentescoID)
	tpGroup.Post("/", controllers.PostTipoParentesco)
	tpGroup.Put("/", controllers.PutTipoParentesco)
	tpGroup.Delete("/:id", controllers.DeleteTipoParentesco)
}