package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var justificacionAusenciaCache = caching.NewCache("JustificacionAusencia", 100)
var justificacionAusenciaSubject = observer.NewSubject()

func init() {
	justificacionAusenciaSubject.AddObserver(updateJustificacionAusenciaCache)
}

func updateJustificacionAusenciaCache(url string) {
	justificacionAusenciaCache.Delete(url)
}

func GetJustificacionesAusencia(c *fiber.Ctx) error {
	if result, exists := justificacionAusenciaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Justificaciones obtenidas del cache", result)
	}

	data, err := handlers.Get(c, models.GetJustificacionesAusencia)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener justificaciones", err)
	}

	justificacionAusenciaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de justificaciones obtenida", data)
}

func GetJustificacionAusenciaID(c *fiber.Ctx) error {
	if result, exists := justificacionAusenciaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Justificación obtenida del cache", result)
	}

	data, err := handlers.GetID(c, models.GetJustificacionAusenciaID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener justificación", err)
	}

	justificacionAusenciaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Justificación obtenida con éxito", data)
}

func PostJustificacionAusencia(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de justificación inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostJustificacionAusencia)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar justificación", err)
	}

	justificacionAusenciaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Justificación registrada exitosamente", nil)
}

func PutJustificacionAusencia(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de justificación inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutJustificacionAusencia)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar justificación", err)
	}

	justificacionAusenciaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Justificación actualizada exitosamente", nil)
}

func DeleteJustificacionAusencia(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteJustificacionAusencia)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar justificación", err)
	}

	justificacionAusenciaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Justificación eliminada exitosamente", nil)
}