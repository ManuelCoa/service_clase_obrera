package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var gestionParentescoCache = caching.NewCache("GestionParentesco", 100)
var gestionParentescoSubject = observer.NewSubject()

func init() {
	gestionParentescoSubject.AddObserver(updateGestionParentescoCache)
}

func updateGestionParentescoCache(url string) {
	gestionParentescoCache.Delete(url)
}

func GetGestionParentescos(c *fiber.Ctx) error {
	if result, exists := gestionParentescoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Parentescos obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetGestionParentescos)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener parentescos", err)
	}

	gestionParentescoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de parentescos obtenida", data)
}

func GetGestionParentescoID(c *fiber.Ctx) error {
	if result, exists := gestionParentescoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Parentesco obtenido del cache", result)
	}

	data, err := handlers.GetID(c, models.GetGestionParentescoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener parentesco", err)
	}

	gestionParentescoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Parentesco obtenido con éxito", data)
}

func PostGestionParentesco(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de parentesco inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostGestionParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar parentesco", err)
	}

	gestionParentescoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Parentesco registrado exitosamente", nil)
}

func PutGestionParentesco(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de parentesco inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutGestionParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar parentesco", err)
	}

	gestionParentescoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Parentesco actualizado exitosamente", nil)
}

func DeleteGestionParentesco(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteGestionParentesco)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar parentesco", err)
	}

	gestionParentescoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Parentesco eliminado exitosamente", nil)
}