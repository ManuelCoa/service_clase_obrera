package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var responsabilidadComunalCache = caching.NewCache("ResponsabilidadComunal", 100)
var responsabilidadComunalSubject = observer.NewSubject()

func init() {
	responsabilidadComunalSubject.AddObserver(updateResponsabilidadComunalCache)
}

func updateResponsabilidadComunalCache(url string) {
	responsabilidadComunalCache.Delete(url)
}

func GetResponsabilidadComunal(c *fiber.Ctx) error {
	if result, exists := responsabilidadComunalCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Responsabilidades comunales obtenidas del caché", result)
	}

	data, err := handlers.Get(c, models.GetResponsabilidadComunal)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener responsabilidades comunales", err)
	}

	responsabilidadComunalCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Responsabilidades comunales obtenidas exitosamente", data)
}

func GetResponsabilidadComunalID(c *fiber.Ctx) error {
	if result, exists := responsabilidadComunalCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Responsabilidad comunal obtenida del caché", result)
	}

	data, err := handlers.GetID(c, models.GetResponsabilidadComunalID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener responsabilidad comunal", err)
	}

	responsabilidadComunalCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Responsabilidad comunal obtenida exitosamente", data)
}

func PostResponsabilidadComunal(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de responsabilidad comunal no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear responsabilidades comunales", err)
	}

	err = handlers.HandlePost(c, models.PostResponsabilidadComunal)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear responsabilidad comunal", err)
	}

	responsabilidadComunalSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Responsabilidad comunal creada exitosamente", nil)
}

func PutResponsabilidadComunal(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de responsabilidad comunal no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar responsabilidades comunales", err)
	}

	err = handlers.Put(c, models.PutResponsabilidadComunal)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar responsabilidad comunal", err)
	}

	responsabilidadComunalSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Responsabilidad comunal actualizada exitosamente", nil)
}

func DeleteResponsabilidadComunal(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteResponsabilidadComunal)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar responsabilidad comunal", err)
	}

	responsabilidadComunalSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Responsabilidad comunal eliminada exitosamente", nil)
}