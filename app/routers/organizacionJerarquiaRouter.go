package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func OrganizacionJerarquiaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	orgGroup := api.Group("/organizacionjerarquia")
	orgGroup.Use(middleware.CorsMiddleware(corsOptions))
	orgGroup.Use(middleware.AuthRequired)
	orgGroup.Use(middleware.CheckPermissions)

	orgGroup.Get("/", controllers.GetOrganizacionJerarquia)
	orgGroup.Get("/:id", controllers.GetOrganizacionJerarquiaID)
	orgGroup.Post("/", controllers.PostOrganizacionJerarquia)
	orgGroup.Put("/", controllers.PutOrganizacionJerarquia)
	orgGroup.Delete("/:id", controllers.DeleteOrganizacionJerarquia)
}
