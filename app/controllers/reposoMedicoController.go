package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var reposoMedicoCache = caching.NewCache("ReposoMedico", 100)
var reposoMedicoSubject = observer.NewSubject()

func init() {
	reposoMedicoSubject.AddObserver(updateReposoMedicoCache)
}

func updateReposoMedicoCache(url string) {
	reposoMedicoCache.Delete(url)
}

func GetReposoMedico(c *fiber.Ctx) error {
	if result, exists := reposoMedicoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Reposos médicos obtenidos del caché", result)
	}

	data, err := handlers.Get(c, models.GetReposoMedico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener reposos médicos", err)
	}

	reposoMedicoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Reposos médicos obtenidos exitosamente", data)
}

func GetReposoMedicoID(c *fiber.Ctx) error {
	if result, exists := reposoMedicoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Reposo médico obtenido del caché", result)
	}

	data, err := handlers.GetID(c, models.GetReposoMedicoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener reposo médico", err)
	}

	reposoMedicoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Reposo médico obtenido exitosamente", data)
}

func PostReposoMedico(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de reposo médico no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para crear reposos médicos", err)
	}

	err = handlers.HandlePost(c, models.PostReposoMedico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear reposo médico", err)
	}

	reposoMedicoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Reposo médico creado exitosamente", nil)
}

func PutReposoMedico(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de reposo médico no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar reposos médicos", err)
	}

	err = handlers.Put(c, models.PutReposoMedico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar reposo médico", err)
	}

	reposoMedicoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Reposo médico actualizado exitosamente", nil)
}

func DeleteReposoMedico(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteReposoMedico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar reposo médico", err)
	}

	reposoMedicoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Reposo médico eliminado exitosamente", nil)
}