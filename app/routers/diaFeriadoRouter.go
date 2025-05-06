package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func DiaFeriadoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	diaFeriadoGroup := api.Group("/dia-feriado")

	diaFeriadoGroup.Use(middleware.CorsMiddleware(corsOptions))
	diaFeriadoGroup.Use(middleware.AuthRequired)
	diaFeriadoGroup.Use(middleware.CheckPermissions)

	diaFeriadoGroup.Get("/", controllers.GetDiasFeriados)
	diaFeriadoGroup.Get("/:id", controllers.GetDiaFeriadoID)
	diaFeriadoGroup.Post("/", controllers.PostDiaFeriado)
	diaFeriadoGroup.Put("/", controllers.PutDiaFeriado)
	diaFeriadoGroup.Delete("/:id", controllers.DeleteDiaFeriado)
}