package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var permisoAsistenciaCache = caching.NewCache("PermisoAsistencia", 120)
var permisoAsistenciaSubject = observer.NewSubject()

func init() {
	permisoAsistenciaSubject.AddObserver(updatePermisoAsistenciaCache)
}

func updatePermisoAsistenciaCache(url string) {
	permisoAsistenciaCache.Delete(url)
}

func GetPermisosAsistencia(c *fiber.Ctx) error {
	if result, exists := permisoAsistenciaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Permisos de asistencia obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetPermisosAsistencia)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener permisos de asistencia", err)
	}

	permisoAsistenciaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de permisos de asistencia obtenida", data)
}

func GetPermisoAsistenciaID(c *fiber.Ctx) error {
	if result, exists := permisoAsistenciaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Permiso de asistencia obtenido del cache", result)
	}

	data, err := handlers.GetID(c, models.GetPermisoAsistenciaID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener permiso de asistencia", err)
	}

	permisoAsistenciaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Permiso de asistencia obtenido con éxito", data)
}

func PostPermisoAsistencia(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de permiso de asistencia inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para registrar permisos", err)
	}

	err = handlers.HandlePost(c, models.PostPermisoAsistencia)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar permiso de asistencia", err)
	}

	permisoAsistenciaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Permiso de asistencia registrado exitosamente", nil)
}

func PutPermisoAsistencia(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de permiso de asistencia inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para actualizar permisos", err)
	}

	err = handlers.Put(c, models.PutPermisoAsistencia)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar permiso de asistencia", err)
	}

	permisoAsistenciaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Permiso de asistencia actualizado exitosamente", nil)
}

func DeletePermisoAsistencia(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeletePermisoAsistencia)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar permiso de asistencia", err)
	}

	permisoAsistenciaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Permiso de asistencia eliminado exitosamente", nil)
}