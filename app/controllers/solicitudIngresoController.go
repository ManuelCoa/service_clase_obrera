package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var solicitudIngresoCache = caching.NewCache("SolicitudIngreso", 100)
var solicitudIngresoSubject = observer.NewSubject()

func init() {
	solicitudIngresoSubject.AddObserver(updateSolicitudIngresoCache)
}

func updateSolicitudIngresoCache(url string) {
	solicitudIngresoCache.Delete(url)
}

func GetSolicitudesIngreso(c *fiber.Ctx) error {
	if result, exists := solicitudIngresoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Solicitudes de ingreso obtenidas del cache", result)
	}

	data, err := handlers.Get(c, models.GetSolicitudesIngreso)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener solicitudes de ingreso", err)
	}

	solicitudIngresoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de solicitudes de ingreso obtenida", data)
}

func GetSolicitudIngresoID(c *fiber.Ctx) error {
	if result, exists := solicitudIngresoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Solicitud de ingreso obtenida del cache", result)
	}

	data, err := handlers.GetID(c, models.GetSolicitudIngresoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener solicitud de ingreso", err)
	}

	solicitudIngresoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Solicitud de ingreso obtenida con éxito", data)
}

func PostSolicitudIngreso(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de solicitud de ingreso inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para registrar solicitudes", err)
	}

	err = handlers.HandlePost(c, models.PostSolicitudIngreso)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar solicitud de ingreso", err)
	}

	solicitudIngresoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Solicitud de ingreso registrada exitosamente", nil)
}

func PutSolicitudIngreso(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de solicitud de ingreso inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para actualizar solicitudes", err)
	}

	err = handlers.Put(c, models.PutSolicitudIngreso)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar solicitud de ingreso", err)
	}

	solicitudIngresoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Solicitud de ingreso actualizada exitosamente", nil)
}

func DeleteSolicitudIngreso(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteSolicitudIngreso)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar solicitud de ingreso", err)
	}

	solicitudIngresoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Solicitud de ingreso eliminada exitosamente", nil)
}