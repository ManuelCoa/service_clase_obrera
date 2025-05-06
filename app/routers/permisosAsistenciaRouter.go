package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PermisoAsistenciaRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}
	permisoAsistenciaGroup := api.Group("/permiso-asistencia")

	permisoAsistenciaGroup.Use(middleware.CorsMiddleware(corsOptions))
	permisoAsistenciaGroup.Use(middleware.AuthRequired)
	permisoAsistenciaGroup.Use(middleware.CheckPermissions)

	permisoAsistenciaGroup.Get("/", controllers.GetPermisosAsistencia)
	permisoAsistenciaGroup.Get("/:id", controllers.GetPermisoAsistenciaID)
	permisoAsistenciaGroup.Post("/", controllers.PostPermisoAsistencia)
	permisoAsistenciaGroup.Put("/", controllers.PutPermisoAsistencia)
	permisoAsistenciaGroup.Delete("/:id", controllers.DeletePermisoAsistencia)
}