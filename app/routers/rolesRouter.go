package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func RolesRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOpcion := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	rolesGroup := api.Group("/roles")

	rolesGroup.Use(middleware.CorsMiddleware(corsOpcion))
	rolesGroup.Use(middleware.AuthRequired)
	rolesGroup.Use(middleware.CheckPermissions)

	rolesGroup.Get("/", controllers.GetRoles)
	rolesGroup.Get("/:id", controllers.GetRolesID)
	rolesGroup.Post("/", controllers.PostRoles)
	rolesGroup.Put("/", controllers.PutRoles)
	rolesGroup.Delete("/:id", controllers.DeleteRoles)
}
