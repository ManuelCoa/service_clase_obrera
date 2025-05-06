package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var tipoEnfermedadCache = caching.NewCache("TipoEnfermedad", 100)
var tipoEnfermedadSubject = observer.NewSubject()

func init() {
	tipoEnfermedadSubject.AddObserver(updateTipoEnfermedadCache)
}

func updateTipoEnfermedadCache(url string) {
	tipoEnfermedadCache.Delete(url)
}

func GetTipoEnfermedad(c *fiber.Ctx) error {
	if result, exists := tipoEnfermedadCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de enfermedad obtenidos del caché", result)
	}

	data, err := handlers.Get(c, models.GetTipoEnfermedad)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipos de enfermedad", err)
	}

	tipoEnfermedadCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de enfermedad obtenidos exitosamente", data)
}

func GetTipoEnfermedadID(c *fiber.Ctx) error {
	if result, exists := tipoEnfermedadCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de enfermedad obtenido del caché", result)
	}

	data, err := handlers.GetID(c, models.GetTipoEnfermedadID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipo de enfermedad", err)
	}

	tipoEnfermedadCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de enfermedad obtenido exitosamente", data)
}

func PostTipoEnfermedad(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de enfermedad no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear tipos de enfermedad", err)
	}

	err = handlers.HandlePost(c, models.PostTipoEnfermedad)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear tipo de enfermedad", err)
	}

	tipoEnfermedadSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Tipo de enfermedad creado exitosamente", nil)
}

func PutTipoEnfermedad(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de enfermedad no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar tipos de enfermedad", err)
	}

	err = handlers.Put(c, models.PutTipoEnfermedad)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar tipo de enfermedad", err)
	}

	tipoEnfermedadSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de enfermedad actualizado exitosamente", nil)
}

func DeleteTipoEnfermedad(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteTipoEnfermedad)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar tipo de enfermedad", err)
	}

	tipoEnfermedadSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de enfermedad eliminado exitosamente", nil)
}