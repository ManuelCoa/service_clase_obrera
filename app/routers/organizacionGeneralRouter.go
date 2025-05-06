package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func OrganizacionGeneralRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	orgGroup := api.Group("/organizaciongeneral")
	orgGroup.Use(middleware.CorsMiddleware(corsOptions))
	orgGroup.Use(middleware.AuthRequired)
	orgGroup.Use(middleware.CheckPermissions)

	orgGroup.Get("/", controllers.GetOrganizacionGeneral)
	orgGroup.Get("/:id", controllers.GetOrganizacionGeneralID)
	orgGroup.Post("/", controllers.PostOrganizacionGeneral)
	orgGroup.Put("/", controllers.PutOrganizacionGeneral)
	orgGroup.Delete("/:id", controllers.DeleteOrganizacionGeneral)
}
