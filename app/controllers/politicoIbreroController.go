package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var politicoObreroCache = caching.NewCache("PoliticoObrero", 100)
var politicoObreroSubject = observer.NewSubject()

func init() {
	politicoObreroSubject.AddObserver(updatePoliticoObreroCache)
}

func updatePoliticoObreroCache(url string) {
	politicoObreroCache.Delete(url)
}

func GetPoliticosObreros(c *fiber.Ctx) error {
	if result, exists := politicoObreroCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Políticos obreros obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetPoliticosObreros)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener políticos obreros", err)
	}

	politicoObreroCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de políticos obreros obtenida", data)
}

func GetPoliticoObreroID(c *fiber.Ctx) error {
	if result, exists := politicoObreroCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Político obrero obtenido del cache", result)
	}

	data, err := handlers.GetID(c, models.GetPoliticoObreroID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener político obrero", err)
	}

	politicoObreroCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Político obrero obtenido con éxito", data)
}

func PostPoliticoObrero(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de político obrero inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostPoliticoObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar político obrero", err)
	}

	politicoObreroSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Político obrero registrado exitosamente", nil)
}

func PutPoliticoObrero(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de político obrero inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutPoliticoObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar político obrero", err)
	}

	politicoObreroSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Político obrero actualizado exitosamente", nil)
}

func DeletePoliticoObrero(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeletePoliticoObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar político obrero", err)
	}

	politicoObreroSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Político obrero eliminado exitosamente", nil)
}