package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TransportePublicoRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST","PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	transportePublicoGroup := api.Group("/transporte-publico")

	transportePublicoGroup.Use(middleware.CorsMiddleware(corsOptions))
	transportePublicoGroup.Use(middleware.AuthRequired)
	transportePublicoGroup.Use(middleware.CheckPermissions)

	transportePublicoGroup.Get("/", controllers.GetTransportePublico)
	transportePublicoGroup.Get("/:id", controllers.GetTransportePublicoID)
	transportePublicoGroup.Post("/", controllers.PostTransportePublico)
	transportePublicoGroup.Put("/", controllers.PutTransportePublico)
	transportePublicoGroup.Delete("/:id", controllers.DeleteTransportePublico)
}