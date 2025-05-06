package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PermisosRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	permisosGroup := api.Group("/permisos")

	permisosGroup.Use(middleware.CorsMiddleware(corsOptions))
	permisosGroup.Use(middleware.AuthRequired)
	permisosGroup.Use(middleware.CheckPermissions)

	permisosGroup.Get("/", controllers.GetPermisos)
	permisosGroup.Get("/:id", controllers.GetPermisosID)
	permisosGroup.Post("/", controllers.PostPermisos)
	permisosGroup.Put("/", controllers.PutPermisos)
	permisosGroup.Delete("/:id", controllers.DeletePermisos)
}
