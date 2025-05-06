package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var saludParentescoCache = caching.NewCache("SaludParentesco", 120)
var saludParentescoSubject = observer.NewSubject()

func init() {
	saludParentescoSubject.AddObserver(updateSaludParentescoCache)
}

func updateSaludParentescoCache(url string) {
	saludParentescoCache.Delete(url)
}

func GetSaludParentescos(c *fiber.Ctx) error {
	if result, exists := saludParentescoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registros de salud parentesco obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetSaludParentescos)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener registros de salud parentesco", err)
	}

	saludParentescoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de salud parentesco obtenida", data)
}

func GetSaludParentescoID(c *fiber.Ctx) error {
	if result, exists := saludParentescoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registro de salud parentesco obtenido del cache", result)
	}

	data, err := handlers.GetID(c, models.GetSaludParentescoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener registro de salud parentesco", err)
	}

	saludParentescoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registro de salud parentesco obtenido", data)
}

func PostSaludParentesco(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de salud parentesco inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para registrar salud parentesco", err)
	}

	err = handlers.HandlePost(c, models.PostSaludParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar salud parentesco", err)
	}

	saludParentescoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Registro de salud parentesco creado exitosamente", nil)
}

func PutSaludParentesco(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de salud parentesco inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para actualizar salud parentesco", err)
	}

	err = handlers.Put(c, models.PutSaludParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar registro de salud parentesco", err)
	}

	saludParentescoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registro de salud parentesco actualizado exitosamente", nil)
}

func DeleteSaludParentesco(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteSaludParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar registro de salud parentesco", err)
	}

	saludParentescoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registro de salud parentesco eliminado exitosamente", nil)
}