package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	usersGroup := api.Group("/users")

	usersGroup.Use(middleware.CorsMiddleware(corsOptions))

	usersGroup.Get("/", controllers.GetUsers)
	usersGroup.Get("/:id", controllers.GetUsersID)
	usersGroup.Post("/", controllers.PostUsers)
	usersGroup.Put("/", controllers.PutUsers)
	usersGroup.Delete("/:id", controllers.DeleteUsers)
}
