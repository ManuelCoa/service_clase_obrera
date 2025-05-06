package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func DepartamentosRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	departamentosGroup := api.Group("/departamentos")

	departamentosGroup.Use(middleware.CorsMiddleware(corsOptions))
	departamentosGroup.Use(middleware.AuthRequired)
	departamentosGroup.Use(middleware.CheckPermissions)

	departamentosGroup.Get("/", controllers.GetDepartamentos)
	departamentosGroup.Get("/:id", controllers.GetDepartamentosID)
	departamentosGroup.Post("/", controllers.PostDepartamentos)
	departamentosGroup.Put("/", controllers.PutDepartamentos)
	departamentosGroup.Delete("/:id", controllers.DeleteDepartamentos)
}