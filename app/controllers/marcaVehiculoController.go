package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var marcaVehiculoCache = caching.NewCache("MarcaVehiculo", 100)
var marcaVehiculoSubject = observer.NewSubject()

func init() {
	marcaVehiculoSubject.AddObserver(updateMarcaVehiculoCache)
}

func updateMarcaVehiculoCache(url string) {
	marcaVehiculoCache.Delete(url)
}

func GetMarcaVehiculo(c *fiber.Ctx) error {
	if result, exists := marcaVehiculoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Marcas de vehículos obtenidas del caché", result)
	}

	data, err := handlers.Get(c, models.GetMarcaVehiculo)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener marcas de vehículos", err)
	}

	marcaVehiculoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Marcas de vehículos obtenidas exitosamente", data)
}

func GetMarcaVehiculoID(c *fiber.Ctx) error {
	if result, exists := marcaVehiculoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Marca de vehículo obtenida del caché", result)
	}

	data, err := handlers.GetID(c, models.GetMarcaVehiculoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener marca de vehículo", err)
	}

	marcaVehiculoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Marca de vehículo obtenida exitosamente", data)
}

func PostMarcaVehiculo(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de marca de vehículo no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear marcas de vehículos", err)
	}

	err = handlers.HandlePost(c, models.PostMarcaVehiculo)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear marca de vehículo", err)
	}

	marcaVehiculoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Marca de vehículo creada exitosamente", nil)
}

func PutMarcaVehiculo(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de marca de vehículo no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar marcas de vehículos", err)
	}

	err = handlers.Put(c, models.PutMarcaVehiculo)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar marca de vehículo", err)
	}

	marcaVehiculoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Marca de vehículo actualizada exitosamente", nil)
}

func DeleteMarcaVehiculo(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteMarcaVehiculo)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar marca de vehículo", err)
	}

	marcaVehiculoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Marca de vehículo eliminada exitosamente", nil)
}