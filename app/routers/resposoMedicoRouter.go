package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ReposoMedicoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	rmGroup := api.Group("/reposo-medico")

	rmGroup.Use(middleware.CorsMiddleware(corsOptions))
	rmGroup.Use(middleware.AuthRequired)
	rmGroup.Use(middleware.CheckPermissions)

	rmGroup.Get("/", controllers.GetReposoMedico)
	rmGroup.Get("/:id", controllers.GetReposoMedicoID)
	rmGroup.Post("/", controllers.PostReposoMedico)
	rmGroup.Put("/", controllers.PutReposoMedico)
	rmGroup.Delete("/:id", controllers.DeleteReposoMedico)
}