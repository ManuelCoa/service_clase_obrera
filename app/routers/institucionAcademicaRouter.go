package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func InstitucionAcademicaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	institucionAcademicaGroup := api.Group("/institucion-academica")

	institucionAcademicaGroup.Use(middleware.CorsMiddleware(corsOptions))
	institucionAcademicaGroup.Use(middleware.AuthRequired)
	institucionAcademicaGroup.Use(middleware.CheckPermissions)

	institucionAcademicaGroup.Get("/", controllers.GetInstitucionesAcademicas)
	institucionAcademicaGroup.Get("/:id", controllers.GetInstitucionAcademicaID)
	institucionAcademicaGroup.Post("/", controllers.PostInstitucionAcademica)
	institucionAcademicaGroup.Put("/", controllers.PutInstitucionAcademica)
	institucionAcademicaGroup.Delete("/:id", controllers.DeleteInstitucionAcademica)
}