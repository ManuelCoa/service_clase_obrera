package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var habilidadesCache = caching.NewCache("Habilidades", 100)
var habilidadesSubject = observer.NewSubject()

func init() {
	habilidadesSubject.AddObserver(updateHabilidadesCache)
}

func updateHabilidadesCache(url string) {
	habilidadesCache.Delete(url)
}

func GetHabilidades(c *fiber.Ctx) error {
	if result, exists := habilidadesCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Habilidades obtenidas del cache", result)
	}

	data, err := handlers.Get(c, models.GetHabilidades)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener habilidades", err)
	}

	habilidadesCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de habilidades obtenida", data)
}

func GetHabilidadID(c *fiber.Ctx) error {
	if result, exists := habilidadesCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Habilidad obtenida del cache", result)
	}

	data, err := handlers.GetID(c, models.GetHabilidadID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener habilidad", err)
	}

	habilidadesCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Habilidad obtenida con éxito", data)
}

func PostHabilidad(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de habilidad inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostHabilidad)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar habilidad", err)
	}

	habilidadesSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Habilidad registrada exitosamente", nil)
}

func PutHabilidad(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de habilidad inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutHabilidad)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar habilidad", err)
	}

	habilidadesSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Habilidad actualizada exitosamente", nil)
}

func DeleteHabilidad(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteHabilidad)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar habilidad", err)
	}

	habilidadesSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Habilidad eliminada exitosamente", nil)
}