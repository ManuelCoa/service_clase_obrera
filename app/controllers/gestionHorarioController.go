package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var gestionHorarioCache = caching.NewCache("GestionHorario", 150)
var gestionHorarioSubject = observer.NewSubject()

func init() {
	gestionHorarioSubject.AddObserver(updateGestionHorarioCache)
}

func updateGestionHorarioCache(url string) {
	gestionHorarioCache.Delete(url)
}

func GetGestionHorarios(c *fiber.Ctx) error {
	if result, exists := gestionHorarioCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Horarios obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetGestionHorarios)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener horarios", err)
	}

	gestionHorarioCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de horarios obtenida", data)
}

func GetGestionHorarioID(c *fiber.Ctx) error {
	if result, exists := gestionHorarioCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Horario obtenido del cache", result)
	}

	data, err := handlers.GetID(c, models.GetGestionHorarioID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener horario", err)
	}

	gestionHorarioCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Horario obtenido con éxito", data)
}

func PostGestionHorario(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de horario inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostGestionHorario)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar horario", err)
	}

	gestionHorarioSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Horario registrado exitosamente", nil)
}

func PutGestionHorario(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de horario inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutGestionHorario)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar horario", err)
	}

	gestionHorarioSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Horario actualizado exitosamente", nil)
}

func DeleteGestionHorario(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteGestionHorario)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar horario", err)
	}

	gestionHorarioSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Horario eliminado exitosamente", nil)
}