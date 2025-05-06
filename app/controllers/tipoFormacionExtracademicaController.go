package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var tfeCache = caching.NewCache("TipoFormacionExtracademica", 100)
var tfeSubject = observer.NewSubject()

func init() {
	tfeSubject.AddObserver(updateTFECache)
}

func updateTFECache(url string) {
	tfeCache.Delete(url)
}

func GetTipoFormacionExtracademica(c *fiber.Ctx) error {
	if result, exists := tfeCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de formación obtenidos del caché", result)
	}

	data, err := handlers.Get(c, models.GetTipoFormacionExtracademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipos de formación", err)
	}

	tfeCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de formación obtenidos exitosamente", data)
}

func GetTipoFormacionExtracademicaID(c *fiber.Ctx) error {
	if result, exists := tfeCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de formación obtenido del caché", result)
	}

	data, err := handlers.GetID(c, models.GetTipoFormacionExtracademicaID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipo de formación", err)
	}

	tfeCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de formación obtenido exitosamente", data)
}

func PostTipoFormacionExtracademica(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de formación no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear tipos de formación", err)
	}

	err = handlers.HandlePost(c, models.PostTipoFormacionExtracademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear tipo de formación", err)
	}

	tfeSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Tipo de formación creado exitosamente", nil)
}

func PutTipoFormacionExtracademica(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de formación no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar tipos de formación", err)
	}

	err = handlers.Put(c, models.PutTipoFormacionExtracademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar tipo de formación", err)
	}

	tfeSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de formación actualizado exitosamente", nil)
}

func DeleteTipoFormacionExtracademica(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteTipoFormacionExtracademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar tipo de formación", err)
	}

	tfeSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de formación eliminado exitosamente", nil)
}