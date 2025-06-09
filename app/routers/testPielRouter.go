package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TestPielRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	testPielGroup := api.Group("/test-piel")

	// Se aplica el middleware de CORS
	testPielGroup.Use(middleware.CorsMiddleware(corsOptions))
	testPielGroup.Use(middleware.AuthRequired)
	testPielGroup.Use(middleware.CheckPermissions)

	testPielGroup.Get("/", controllers.GetTestPiel)
	testPielGroup.Get("/:id", controllers.GetTestPielID)
	testPielGroup.Post("/", controllers.PostTestPiel)
	testPielGroup.Put("/", controllers.PutTestPiel)
	testPielGroup.Delete("/:id", controllers.DeleteTestPiel)
}