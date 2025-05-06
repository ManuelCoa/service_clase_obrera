package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var ttCache = caching.NewCache("TipoTecho", 100)
var ttSubject = observer.NewSubject()

func init() {
	ttSubject.AddObserver(updateTTCache)
}

func updateTTCache(url string) {
	ttCache.Delete(url)
}

func GetTipoTecho(c *fiber.Ctx) error {
	if result, exists := ttCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de techo obtenidos del caché", result)
	}

	data, err := handlers.Get(c, models.GetTipoTecho)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipos de techo", err)
	}

	ttCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de techo obtenidos exitosamente", data)
}

func GetTipoTechoID(c *fiber.Ctx) error {
	if result, exists := ttCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de techo obtenido del caché", result)
	}

	data, err := handlers.GetID(c, models.GetTipoTechoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipo de techo", err)
	}

	ttCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de techo obtenido exitosamente", data)
}

func PostTipoTecho(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de techo no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear tipos de techo", err)
	}

	err = handlers.HandlePost(c, models.PostTipoTecho)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear tipo de techo", err)
	}

	ttSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Tipo de techo creado exitosamente", nil)
}

func PutTipoTecho(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de techo no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar tipos de techo", err)
	}

	err = handlers.Put(c, models.PutTipoTecho)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar tipo de techo", err)
	}

	ttSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de techo actualizado exitosamente", nil)
}

func DeleteTipoTecho(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteTipoTecho)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar tipo de techo", err)
	}

	ttSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de techo eliminado exitosamente", nil)
}