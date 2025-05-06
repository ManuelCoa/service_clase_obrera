package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var solicitudEgresoCache = caching.NewCache("SolicitudEgreso", 100)
var solicitudEgresoSubject = observer.NewSubject()

func init() {
	solicitudEgresoSubject.AddObserver(updateSolicitudEgresoCache)
}

func updateSolicitudEgresoCache(url string) {
	solicitudEgresoCache.Delete(url)
}

func GetSolicitudesEgreso(c *fiber.Ctx) error {
	if result, exists := solicitudEgresoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Solicitudes de egreso obtenidas del cache", result)
	}

	data, err := handlers.Get(c, models.GetSolicitudesEgreso)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener solicitudes de egreso", err)
	}

	solicitudEgresoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de solicitudes de egreso obtenida", data)
}

func GetSolicitudEgresoID(c *fiber.Ctx) error {
	if result, exists := solicitudEgresoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Solicitud de egreso obtenida del cache", result)
	}

	data, err := handlers.GetID(c, models.GetSolicitudEgresoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener solicitud de egreso", err)
	}

	solicitudEgresoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Solicitud de egreso obtenida con éxito", data)
}

func PostSolicitudEgreso(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de solicitud de egreso inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para registrar solicitudes", err)
	}

	err = handlers.HandlePost(c, models.PostSolicitudEgreso)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar solicitud de egreso", err)
	}

	solicitudEgresoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Solicitud de egreso registrada exitosamente", nil)
}

func PutSolicitudEgreso(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de solicitud de egreso inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para actualizar solicitudes", err)
	}

	err = handlers.Put(c, models.PutSolicitudEgreso)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar solicitud de egreso", err)
	}

	solicitudEgresoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Solicitud de egreso actualizada exitosamente", nil)
}

func DeleteSolicitudEgreso(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteSolicitudEgreso)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar solicitud de egreso", err)
	}

	solicitudEgresoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Solicitud de egreso eliminada exitosamente", nil)
}