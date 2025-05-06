package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var datosAcademicosCache = caching.NewCache("DatosAcademicos", 150)
var datosAcademicosSubject = observer.NewSubject()

func init() {
	datosAcademicosSubject.AddObserver(updateDatosAcademicosCache)
}

func updateDatosAcademicosCache(url string) {
	datosAcademicosCache.Delete(url)
}

func GetDatosAcademicos(c *fiber.Ctx) error {
	if result, exists := datosAcademicosCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos académicos obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetDatosAcademicos)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener datos académicos", err)
	}

	datosAcademicosCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos académicos obtenidos", data)
}

func GetDatosAcademicosID(c *fiber.Ctx) error {
	if result, exists := datosAcademicosCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos académicos obtenidos del cache", result)
	}

	data, err := handlers.GetID(c, models.GetDatosAcademicosID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener datos académicos", err)
	}

	datosAcademicosCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos académicos obtenidos", data)
}

func GetDatosAcademicosPorObrero(c *fiber.Ctx) error {
	if result, exists := datosAcademicosCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos académicos por obrero obtenidos del cache", result)
	}

	data, err := handlers.GetID(c, models.GetDatosAcademicosPorObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener datos académicos por obrero", err)
	}

	datosAcademicosCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos académicos por obrero obtenidos", data)
}

func PostDatosAcademicos(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos académicos inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostDatosAcademicos)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar datos académicos", err)
	}

	datosAcademicosSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Datos académicos registrados", nil)
}

func PutDatosAcademicos(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos académicos inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutDatosAcademicos)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar datos académicos", err)
	}

	datosAcademicosSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos académicos actualizados", nil)
}

func DeleteDatosAcademicos(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteDatosAcademicos)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar datos académicos", err)
	}

	datosAcademicosSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos académicos eliminados", nil)
}