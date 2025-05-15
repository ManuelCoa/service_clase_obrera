package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var direccionParentescocache = caching.NewCache("Knapsack", 100)
var direccionParentescoSubject = observer.NewSubject()

func init() {
	direccionParentescoSubject.AddObserver(updateDireccionParentescoCache)
}

// updateDireccionParentescoCache es una función observadora que actualiza la caché
func updateDireccionParentescoCache(url string) {
	direccionParentescocache.Delete(url)
}

func GetDireccionParentesco(c *fiber.Ctx) error {
	if result, exists := direccionParentescocache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	data, err := handlers.Get(c, models.GetDireccionParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	direccionParentescocache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func GetDireccionParentescoID(c *fiber.Ctx) error {
	if result, exists := direccionParentescocache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	data, err := handlers.GetID(c, models.GetDireccionParentescoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	direccionParentescocache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func PostDireccionParentesco(c *fiber.Ctx) error {
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

	err = handlers.HandlePost(c, models.PostDireccionParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	direccionParentescoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "La dirección parentesco se ha creado con éxito", nil)
}

func PutDireccionParentesco(c *fiber.Ctx) error {
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

	err = handlers.Put(c, models.PutDireccionParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	direccionParentescoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "La dirección parentesco se ha actualizado con éxito", nil)
}

func DeleteDireccionParentesco(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteDireccionParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	direccionParentescoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "La dirección parentesco se ha eliminado con éxito", nil)
}