package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var registroTalentoHumanoCache = caching.NewCache("RegistroTalentoHumano", 150)
var registroTalentoHumanoSubject = observer.NewSubject()

func init() {
	registroTalentoHumanoSubject.AddObserver(updateRegistroTalentoHumanoCache)
}

func updateRegistroTalentoHumanoCache(url string) {
	registroTalentoHumanoCache.Delete(url)
}

func GetRegistrosTalentoHumano(c *fiber.Ctx) error {
	if result, exists := registroTalentoHumanoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registros de talento humano obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetRegistrosTalentoHumano)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener registros de talento humano", err)
	}

	registroTalentoHumanoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de registros de talento humano obtenida", data)
}

func GetRegistroTalentoHumanoID(c *fiber.Ctx) error {
	if result, exists := registroTalentoHumanoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registro de talento humano obtenido del cache", result)
	}

	data, err := handlers.GetID(c, models.GetRegistroTalentoHumanoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener registro de talento humano", err)
	}

	registroTalentoHumanoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registro de talento humano obtenido con éxito", data)
}

func PostRegistroTalentoHumano(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de registro de talento humano inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para registrar talento humano", err)
	}

	err = handlers.HandlePost(c, models.PostRegistroTalentoHumano)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar talento humano", err)
	}

	registroTalentoHumanoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Registro de talento humano creado exitosamente", nil)
}

func PutRegistroTalentoHumano(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de registro de talento humano inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para actualizar talento humano", err)
	}

	err = handlers.Put(c, models.PutRegistroTalentoHumano)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar registro de talento humano", err)
	}

	registroTalentoHumanoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registro de talento humano actualizado exitosamente", nil)
}

func DeleteRegistroTalentoHumano(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteRegistroTalentoHumano)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar registro de talento humano", err)
	}

	registroTalentoHumanoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registro de talento humano eliminado exitosamente", nil)
}