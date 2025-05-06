package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var tvCache = caching.NewCache("TipoVivienda", 100)
var tvSubject = observer.NewSubject()

func init() {
	tvSubject.AddObserver(updateTVCache)
}

func updateTVCache(url string) {
	tvCache.Delete(url)
}

func GetTipoVivienda(c *fiber.Ctx) error {
	if result, exists := tvCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de vivienda obtenidos del caché", result)
	}

	data, err := handlers.Get(c, models.GetTipoVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipos de vivienda", err)
	}

	tvCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de vivienda obtenidos exitosamente", data)
}

func GetTipoViviendaID(c *fiber.Ctx) error {
	if result, exists := tvCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de vivienda obtenido del caché", result)
	}

	data, err := handlers.GetID(c, models.GetTipoViviendaID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipo de vivienda", err)
	}

	tvCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de vivienda obtenido exitosamente", data)
}

func PostTipoVivienda(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de vivienda no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear tipos de vivienda", err)
	}

	err = handlers.HandlePost(c, models.PostTipoVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear tipo de vivienda", err)
	}

	tvSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Tipo de vivienda creado exitosamente", nil)
}

func PutTipoVivienda(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de vivienda no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar tipos de vivienda", err)
	}

	err = handlers.Put(c, models.PutTipoVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar tipo de vivienda", err)
	}

	tvSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de vivienda actualizado exitosamente", nil)
}

func DeleteTipoVivienda(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteTipoVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar tipo de vivienda", err)
	}

	tvSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de vivienda eliminado exitosamente", nil)
}