package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var diaFeriadoCache = caching.NewCache("DiaFeriado", 50) // Menor capacidad por ser datos menos volátiles
var diaFeriadoSubject = observer.NewSubject()

func init() {
	diaFeriadoSubject.AddObserver(updateDiaFeriadoCache)
}

func updateDiaFeriadoCache(url string) {
	diaFeriadoCache.Delete(url)
}

func GetDiasFeriados(c *fiber.Ctx) error {
	if result, exists := diaFeriadoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Días feriados obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetDiasFeriados)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener días feriados", err)
	}

	diaFeriadoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de días feriados obtenida", data)
}

func GetDiaFeriadoID(c *fiber.Ctx) error {
	if result, exists := diaFeriadoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Día feriado obtenido del cache", result)
	}

	data, err := handlers.GetID(c, models.GetDiaFeriadoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener día feriado", err)
	}

	diaFeriadoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Día feriado obtenido", data)
}

func PostDiaFeriado(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de día feriado inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostDiaFeriado)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar día feriado", err)
	}

	diaFeriadoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Día feriado registrado exitosamente", nil)
}

func PutDiaFeriado(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de día feriado inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutDiaFeriado)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar día feriado", err)
	}

	diaFeriadoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Día feriado actualizado exitosamente", nil)
}

func DeleteDiaFeriado(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteDiaFeriado)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar día feriado", err)
	}

	diaFeriadoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Día feriado eliminado exitosamente", nil)
}