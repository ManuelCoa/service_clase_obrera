package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func MunicipiosRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	municipioGroup := api.Group("/municipio")
	municipioGroup.Use(middleware.CorsMiddleware(corsOptions))
	municipioGroup.Use(middleware.AuthRequired)
	municipioGroup.Use(middleware.CheckPermissions)

	municipioGroup.Get("/", controllers.GetMunicipio)
	municipioGroup.Get("/:id", controllers.GetMunicipioID)
	municipioGroup.Post("/", controllers.PostMunicipio)
	municipioGroup.Put("/", controllers.PutMunicipio)
	municipioGroup.Delete("/:id", controllers.DeleteMunicipio)
}
