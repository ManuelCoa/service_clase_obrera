package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var tsvCache = caching.NewCache("TipoSueloVivienda", 100)
var tsvSubject = observer.NewSubject()

func init() {
	tsvSubject.AddObserver(updateTSVCache)
}

func updateTSVCache(url string) {
	tsvCache.Delete(url)
}

func GetTipoSueloVivienda(c *fiber.Ctx) error {
	if result, exists := tsvCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de suelo obtenidos del caché", result)
	}

	data, err := handlers.Get(c, models.GetTipoSueloVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipos de suelo", err)
	}

	tsvCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de suelo obtenidos exitosamente", data)
}

func GetTipoSueloViviendaID(c *fiber.Ctx) error {
	if result, exists := tsvCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de suelo obtenido del caché", result)
	}

	data, err := handlers.GetID(c, models.GetTipoSueloViviendaID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipo de suelo", err)
	}

	tsvCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de suelo obtenido exitosamente", data)
}

func PostTipoSueloVivienda(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de suelo no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear tipos de suelo", err)
	}

	err = handlers.HandlePost(c, models.PostTipoSueloVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear tipo de suelo", err)
	}

	tsvSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Tipo de suelo creado exitosamente", nil)
}

func PutTipoSueloVivienda(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de suelo no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar tipos de suelo", err)
	}

	err = handlers.Put(c, models.PutTipoSueloVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar tipo de suelo", err)
	}

	tsvSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de suelo actualizado exitosamente", nil)
}

func DeleteTipoSueloVivienda(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteTipoSueloVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar tipo de suelo", err)
	}

	tsvSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de suelo eliminado exitosamente", nil)
}