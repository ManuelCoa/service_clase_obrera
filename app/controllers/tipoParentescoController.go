package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var tpCache = caching.NewCache("TipoParentesco", 100)
var tpSubject = observer.NewSubject()

func init() {
	tpSubject.AddObserver(updateTPCache)
}

func updateTPCache(url string) {
	tpCache.Delete(url)
}

func GetTipoParentesco(c *fiber.Ctx) error {
	if result, exists := tpCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de parentesco obtenidos del caché", result)
	}

	data, err := handlers.Get(c, models.GetTipoParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipos de parentesco", err)
	}

	tpCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipos de parentesco obtenidos exitosamente", data)
}

func GetTipoParentescoID(c *fiber.Ctx) error {
	if result, exists := tpCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de parentesco obtenido del caché", result)
	}

	data, err := handlers.GetID(c, models.GetTipoParentescoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener tipo de parentesco", err)
	}

	tpCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de parentesco obtenido exitosamente", data)
}

func PostTipoParentesco(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de parentesco no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear tipos de parentesco", err)
	}

	err = handlers.HandlePost(c, models.PostTipoParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear tipo de parentesco", err)
	}

	tpSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Tipo de parentesco creado exitosamente", nil)
}

func PutTipoParentesco(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de tipo de parentesco no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar tipos de parentesco", err)
	}

	err = handlers.Put(c, models.PutTipoParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar tipo de parentesco", err)
	}

	tpSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de parentesco actualizado exitosamente", nil)
}

func DeleteTipoParentesco(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteTipoParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar tipo de parentesco", err)
	}

	tpSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Tipo de parentesco eliminado exitosamente", nil)
}