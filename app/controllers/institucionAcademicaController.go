package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var institucionAcademicaCache = caching.NewCache("InstitucionAcademica", 100)
var institucionAcademicaSubject = observer.NewSubject()

func init() {
	institucionAcademicaSubject.AddObserver(updateInstitucionAcademicaCache)
}

func updateInstitucionAcademicaCache(url string) {
	institucionAcademicaCache.Delete(url)
}

func GetInstitucionesAcademicas(c *fiber.Ctx) error {
	if result, exists := institucionAcademicaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Instituciones académicas obtenidas del cache", result)
	}

	data, err := handlers.Get(c, models.GetInstitucionesAcademicas)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener instituciones académicas", err)
	}

	institucionAcademicaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de instituciones académicas obtenida", data)
}

func GetInstitucionAcademicaID(c *fiber.Ctx) error {
	if result, exists := institucionAcademicaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Institución académica obtenida del cache", result)
	}

	data, err := handlers.GetID(c, models.GetInstitucionAcademicaID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener institución académica", err)
	}

	institucionAcademicaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Institución académica obtenida", data)
}

func PostInstitucionAcademica(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de institución académica inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostInstitucionAcademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar institución académica", err)
	}

	institucionAcademicaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Institución académica registrada exitosamente", nil)
}

func PutInstitucionAcademica(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de institución académica inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutInstitucionAcademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar institución académica", err)
	}

	institucionAcademicaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Institución académica actualizada exitosamente", nil)
}

func DeleteInstitucionAcademica(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteInstitucionAcademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar institución académica", err)
	}

	institucionAcademicaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Institución académica eliminada exitosamente", nil)
}