package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var datosLaboralesCache = caching.NewCache("DatosLaborales", 150)
var datosLaboralesSubject = observer.NewSubject()

func init() {
	datosLaboralesSubject.AddObserver(updateDatosLaboralesCache)
}

func updateDatosLaboralesCache(url string) {
	datosLaboralesCache.Delete(url)
}

func GetDatosLaborales(c *fiber.Ctx) error {
	if result, exists := datosLaboralesCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos laborales obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetDatosLaborales)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener datos laborales", err)
	}

	datosLaboralesCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos laborales obtenidos", data)
}

func GetDatosLaboralesID(c *fiber.Ctx) error {
	if result, exists := datosLaboralesCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos laborales obtenidos del cache", result)
	}

	data, err := handlers.GetID(c, models.GetDatosLaboralesID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener datos laborales", err)
	}

	datosLaboralesCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos laborales obtenidos", data)
}

func GetDatosLaboralesPorObrero(c *fiber.Ctx) error {
	if result, exists := datosLaboralesCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos laborales por obrero obtenidos del cache", result)
	}

	data, err := handlers.GetID(c, models.GetDatosLaboralesPorObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener datos laborales por obrero", err)
	}

	datosLaboralesCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos laborales por obrero obtenidos", data)
}

func PostDatosLaborales(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos laborales inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostDatosLaborales)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar datos laborales", err)
	}

	datosLaboralesSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Datos laborales registrados", nil)
}

func PutDatosLaborales(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos laborales inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutDatosLaborales)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar datos laborales", err)
	}

	datosLaboralesSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos laborales actualizados", nil)
}

func DeleteDatosLaborales(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteDatosLaborales)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar datos laborales", err)
	}

	datosLaboralesSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos laborales eliminados", nil)
}