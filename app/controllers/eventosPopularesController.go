package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var eventoPopularCache = caching.NewCache("Knapsack", 100)
var eventoPopularSubject = observer.NewSubject()

func init() {
	eventoPopularSubject.AddObserver(updateEventoPopularCache)
}

func updateEventoPopularCache(url string) {
	eventoPopularCache.Delete(url)
}

func GetEventoPopular(c *fiber.Ctx) error {
	if result, exists := eventoPopularCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Eventos populares obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetEventoPopular)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener eventos populares", err)
	}

	eventoPopularCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Eventos populares obtenidos con éxito", data)
}

func GetEventoPopularID(c *fiber.Ctx) error {
	if result, exists := eventoPopularCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Evento popular obtenido del cache", result)
	}

	data, err := handlers.GetID(c, models.GetEventoPopularID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener evento popular", err)
	}

	eventoPopularCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Evento popular obtenido con éxito", data)
}

func PostEventoPopular(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de evento popular no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para evento popular", err)
	}

	err = handlers.HandlePost(c, models.PostEventoPopular)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear evento popular", err)
	}

	eventoPopularSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Evento popular creado con éxito", nil)
}

func PutEventoPopular(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de evento popular no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para evento popular", err)
	}

	err = handlers.Put(c, models.PutEventoPopular)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar evento popular", err)
	}

	eventoPopularSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Evento popular actualizado con éxito", nil)
}

func DeleteEventoPopular(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteEventoPopular)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar evento popular", err)
	}

	eventoPopularSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Evento popular eliminado con éxito", nil)
}