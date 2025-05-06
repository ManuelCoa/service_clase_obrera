package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func MisionesBaseRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	mbGroup := api.Group("/misiones-base")

	mbGroup.Use(middleware.CorsMiddleware(corsOptions))
	mbGroup.Use(middleware.AuthRequired)
	mbGroup.Use(middleware.CheckPermissions)

	mbGroup.Get("/", controllers.GetMisionesBase)
	mbGroup.Get("/:id", controllers.GetMisionesBaseID)
	mbGroup.Post("/", controllers.PostMisionesBase)
	mbGroup.Put("/", controllers.PutMisionesBase)
	mbGroup.Delete("/:id", controllers.DeleteMisionesBase)
}