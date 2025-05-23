package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	//"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func DataObreroRoutes(app *fiber.App) {
	api := app.Group("/api")

	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	
	obreroGroup := api.Group("/obrero")

	obreroGroup.Use(middleware.CorsMiddleware(corsOptions))
	obreroGroup.Use(middleware.AuthRequired)
	obreroGroup.Use(middleware.CheckPermissions)
	
	obreroGroup.Get("/", controllers.GetDataObrero)
	obreroGroup.Get("/:cedula", controllers.GetDataObreroID)
	obreroGroup.Get("inst/:id_institucion", controllers.GetDataObreroIDInstitucion)
	obreroGroup.Post("/", controllers.PostDataObrero)
	obreroGroup.Put("/", controllers.PutDataObrero)
	obreroGroup.Delete("/:cedula", controllers.DeleteDataObrero)
}