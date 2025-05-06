package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var misionesBaseCache = caching.NewCache("MisionesBase", 100)
var misionesBaseSubject = observer.NewSubject()

func init() {
	misionesBaseSubject.AddObserver(updateMisionesBaseCache)
}

func updateMisionesBaseCache(url string) {
	misionesBaseCache.Delete(url)
}

func GetMisionesBase(c *fiber.Ctx) error {
	if result, exists := misionesBaseCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Misiones base obtenidas del caché", result)
	}

	data, err := handlers.Get(c, models.GetMisionesBase)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener misiones base", err)
	}

	misionesBaseCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Misiones base obtenidas exitosamente", data)
}

func GetMisionesBaseID(c *fiber.Ctx) error {
	if result, exists := misionesBaseCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Misión base obtenida del caché", result)
	}

	data, err := handlers.GetID(c, models.GetMisionesBaseID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener misión base", err)
	}

	misionesBaseCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Misión base obtenida exitosamente", data)
}

func PostMisionesBase(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de misión base no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear misiones base", err)
	}

	err = handlers.HandlePost(c, models.PostMisionesBase)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear misión base", err)
	}

	misionesBaseSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Misión base creada exitosamente", nil)
}

func PutMisionesBase(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de misión base no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar misiones base", err)
	}

	err = handlers.Put(c, models.PutMisionesBase)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar misión base", err)
	}

	misionesBaseSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Misión base actualizada exitosamente", nil)
}

func DeleteMisionesBase(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteMisionesBase)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar misión base", err)
	}

	misionesBaseSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Misión base eliminada exitosamente", nil)
}