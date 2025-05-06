package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var horarioCache = caching.NewCache("Horario", 100)
var horarioSubject = observer.NewSubject()

func init() {
	horarioSubject.AddObserver(updateHorarioCache)
}

func updateHorarioCache(url string) {
	horarioCache.Delete(url)
}

func GetHorario(c *fiber.Ctx) error {
	if result, exists := horarioCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos de horario obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetHorario)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener horarios", err)
	}

	horarioCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Horarios obtenidos exitosamente", data)
}

func GetHorarioID(c *fiber.Ctx) error {
	if result, exists := horarioCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos de horario obtenidos del cache", result)
	}

	data, err := handlers.GetID(c, models.GetHorarioID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener horario", err)
	}

	horarioCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Horario obtenido exitosamente", data)
}

func PostHorario(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de horario no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear horarios", err)
	}

	err = handlers.HandlePost(c, models.PostHorario)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear horario", err)
	}

	horarioSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Horario creado exitosamente", nil)
}

func PutHorario(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de horario no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar horarios", err)
	}

	err = handlers.Put(c, models.PutHorario)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar horario", err)
	}

	horarioSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Horario actualizado exitosamente", nil)
}

func DeleteHorario(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteHorario)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar horario", err)
	}

	horarioSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Horario eliminado exitosamente", nil)
}