package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AcademicaParentescoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	academicaParentescoGroup := api.Group("/academica-parentesco")

	// Se aplica el middleware de CORS
	academicaParentescoGroup.Use(middleware.CorsMiddleware(corsOptions))
	academicaParentescoGroup.Use(middleware.AuthRequired)
	academicaParentescoGroup.Use(middleware.CheckPermissions)

	academicaParentescoGroup.Get("/", controllers.GetAcademicaParentesco)
	academicaParentescoGroup.Get("/:id", controllers.GetAcademicaParentescoID)
	academicaParentescoGroup.Post("/", controllers.PostAcademicaParentesco)
	academicaParentescoGroup.Put("/", controllers.PutAcademicaParentesco)
	academicaParentescoGroup.Delete("/:id", controllers.DeleteAcademicaParentesco)
}