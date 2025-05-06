package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var bioquimicosObreroCache = caching.NewCache("BioquimicosObrero", 100)
var bioquimicosObreroSubject = observer.NewSubject()

func init() {
	bioquimicosObreroSubject.AddObserver(updateBioquimicosObreroCache)
}

func updateBioquimicosObreroCache(url string) {
	bioquimicosObreroCache.Delete(url)
}

func GetBioquimicosObreros(c *fiber.Ctx) error {
	if result, exists := bioquimicosObreroCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	data, err := handlers.Get(c, models.GetBioquimicosObreros)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	bioquimicosObreroCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func GetBioquimicosObreroID(c *fiber.Ctx) error {
	if result, exists := bioquimicosObreroCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	data, err := handlers.GetID(c, models.GetBioquimicosObreroID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	bioquimicosObreroCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func PostBioquimicosObrero(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de entrada no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostBioquimicosObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	bioquimicosObreroSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Los datos bioquímicos del obrero se han creado con éxito", nil)
}

func PutBioquimicosObrero(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de entrada no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutBioquimicosObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	bioquimicosObreroSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Los datos bioquímicos del obrero se han actualizado con éxito", nil)
}

func DeleteBioquimicosObrero(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteBioquimicosObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	bioquimicosObreroSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Los datos bioquímicos del obrero se han eliminado con éxito", nil)
}