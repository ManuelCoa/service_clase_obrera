package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var institucionExtracademicaCache = caching.NewCache("InstitucionExtracademica", 100)
var institucionExtracademicaSubject = observer.NewSubject()

func init() {
	institucionExtracademicaSubject.AddObserver(updateInstitucionExtracademicaCache)
}

func updateInstitucionExtracademicaCache(url string) {
	institucionExtracademicaCache.Delete(url)
}

func GetInstitucionExtracademica(c *fiber.Ctx) error {
	if result, exists := institucionExtracademicaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos de instituciones obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetInstitucionExtracademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener instituciones", err)
	}

	institucionExtracademicaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Instituciones obtenidas exitosamente", data)
}

func GetInstitucionExtracademicaID(c *fiber.Ctx) error {
	if result, exists := institucionExtracademicaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos de institución obtenidos del cache", result)
	}

	data, err := handlers.GetID(c, models.GetInstitucionExtracademicaID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener institución", err)
	}

	institucionExtracademicaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Institución obtenida exitosamente", data)
}

func PostInstitucionExtracademica(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de institución no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear instituciones", err)
	}

	err = handlers.HandlePost(c, models.PostInstitucionExtracademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear institución", err)
	}

	institucionExtracademicaSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Institución creada exitosamente", nil)
}

func PutInstitucionExtracademica(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de institución no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar instituciones", err)
	}

	err = handlers.Put(c, models.PutInstitucionExtracademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar institución", err)
	}

	institucionExtracademicaSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Institución actualizada exitosamente", nil)
}

func DeleteInstitucionExtracademica(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteInstitucionExtracademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar institución", err)
	}

	institucionExtracademicaSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Institución eliminada exitosamente", nil)
}