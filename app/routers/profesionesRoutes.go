package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProfesionRoutes(app *fiber.App) {
	api := app.Group("/api")
	
		corsOptions := middleware.CorsOptions{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
			AllowedHeaders:   []string{"Content-Type", "Authorization"},
			AllowCredentials: true,
			ExposedHeaders:   []string{"X-Custom-Header"},
			MaxAge:           3600,
		}
	
	profesionesGroup := api.Group("/profesiones")

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	
		profesionesGroup.Use(middleware.CorsMiddleware(corsOptions))
		profesionesGroup.Use(middleware.AuthRequired)
		profesionesGroup.Use(middleware.CheckPermissions)
	

	profesionesGroup.Get("/", controllers.GetProfesiones)
	profesionesGroup.Get("/:id", controllers.GetProfesionesID)
	profesionesGroup.Post("/", controllers.PostProfesiones)
	profesionesGroup.Put("/", controllers.PutProfesiones)
	profesionesGroup.Delete("/:id", controllers.DeleteProfesiones)
}
