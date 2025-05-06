package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var thCache = caching.NewCache("TipoHabilidad", 100)
var thSubject = observer.NewSubject()

func init() {
	thSubject.AddObserver(updateTHCache)
}

func updateTHCache(url string) {
	thCache.Delete(url)
}

func GetTipoHabilidad(c *fiber.Ctx) error {
	if result, exists := thCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de habilidad obtenidos del caché", result)
	}

	data, err := handlers.Get(c, models.GetTipoHabilidad)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipos de habilidad", err)
	}

	thCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de habilidad obtenidos exitosamente", data)
}

func GetTipoHabilidadID(c *fiber.Ctx) error {
	if result, exists := thCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de habilidad obtenido del caché", result)
	}

	data, err := handlers.GetID(c, models.GetTipoHabilidadID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipo de habilidad", err)
	}

	thCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de habilidad obtenido exitosamente", data)
}

func PostTipoHabilidad(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de habilidad no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear tipos de habilidad", err)
	}

	err = handlers.HandlePost(c, models.PostTipoHabilidad)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear tipo de habilidad", err)
	}

	thSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Tipo de habilidad creado exitosamente", nil)
}

func PutTipoHabilidad(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de habilidad no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar tipos de habilidad", err)
	}

	err = handlers.Put(c, models.PutTipoHabilidad)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar tipo de habilidad", err)
	}

	thSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de habilidad actualizado exitosamente", nil)
}

func DeleteTipoHabilidad(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteTipoHabilidad)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar tipo de habilidad", err)
	}

	thSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de habilidad eliminado exitosamente", nil)
}