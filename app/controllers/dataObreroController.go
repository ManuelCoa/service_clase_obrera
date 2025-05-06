package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var obreroCache = caching.NewCache("DataObrero", 150) // Mayor capacidad para datos de obreros
var obreroSubject = observer.NewSubject()

func init() {
	obreroSubject.AddObserver(updateObreroCache)
}

func updateObreroCache(url string) {
	obreroCache.Delete(url)
}

func GetDataObrero(c *fiber.Ctx) error {
	if result, exists := obreroCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos de obreros obtenidos del caché", result)
	}

	data, err := handlers.Get(c, models.GetDataObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener datos de obreros", err)
	}

	obreroCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos de obreros obtenidos exitosamente", data)
}

func GetDataObreroID(c *fiber.Ctx) error {

	if result, exists := obreroCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos del obrero obtenidos del caché", result)
	}

	data, err := handlers.GetID(c, models.GetDataObreroID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener datos del obrero", err)
	}

	obreroCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos del obrero obtenidos exitosamente", data)
}

func PostDataObrero(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos del obrero no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para registrar obreros", err)
	}

	err = handlers.HandlePost(c, models.PostDataObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar obrero", err)
	}

	obreroSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Obrero registrado exitosamente", nil)
}

func PutDataObrero(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos del obrero no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar datos de obreros", err)
	}

	err = handlers.Put(c, models.PutDataObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar datos del obrero", err)
	}

	obreroSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos del obrero actualizados exitosamente", nil)
}

func DeleteDataObrero(c *fiber.Ctx) error {
	
	err := handlers.Delete(c, models.DeleteDataObrero)

	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar obrero", err)
	}

	obreroSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Obrero eliminado exitosamente", nil)
}