package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var gestionExtracademicaCache = caching.NewCache("GestionExtracademica", 100)
var gestionExtracademicaSubject = observer.NewSubject()

func init() {
	gestionExtracademicaSubject.AddObserver(updateGestionExtracademicaCache)
}

func updateGestionExtracademicaCache(url string) {
	gestionExtracademicaCache.Delete(url)
}

func GetGestionesExtracademicas(c *fiber.Ctx) error {
	if result, exists := gestionExtracademicaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Gestiones extracadémicas obtenidas del cache", result)
	}

	data, err := handlers.Get(c, models.GetGestionesExtracademicas)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener gestiones extracadémicas", err)
	}

	gestionExtracademicaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Gestiones extracadémicas obtenidas con éxito", data)
}

func GetGestionExtracademicaID(c *fiber.Ctx) error {
	if result, exists := gestionExtracademicaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Gestión extracadémica obtenida del cache", result)
	}

	data, err := handlers.GetID(c, models.GetGestionExtracademicaID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener gestión extracadémica", err)
	}

	gestionExtracademicaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Gestión extracadémica obtenida con éxito", data)
}

func PostGestionExtracademica(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de gestión extracadémica inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostGestionExtracademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar gestión extracadémica", err)
	}

	gestionExtracademicaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Gestión extracadémica registrada exitosamente", nil)
}

func PutGestionExtracademica(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de gestión extracadémica inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutGestionExtracademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar gestión extracadémica", err)
	}

	gestionExtracademicaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Gestión extracadémica actualizada exitosamente", nil)
}

func DeleteGestionExtracademica(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteGestionExtracademica)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar gestión extracadémica", err)
	}

	gestionExtracademicaSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Gestión extracadémica eliminada exitosamente", nil)
}