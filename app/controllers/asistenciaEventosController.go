package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var asistenciaEventoCache = caching.NewCache("Knapsack", 100)
var asistenciaEventoSubject = observer.NewSubject()

func init() {
	asistenciaEventoSubject.AddObserver(updateAsistenciaEventoCache)
}

// updateAsistenciaEventoCache actualiza la caché cuando hay cambios
func updateAsistenciaEventoCache(url string) {
	asistenciaEventoCache.Delete(url)
}

func GetAsistenciaEvento(c *fiber.Ctx) error {
	// Verificar caché primero
	if result, exists := asistenciaEventoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de asistencias obtenida del caché", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.Get(c, models.GetAsistenciaEvento)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener asistencias", err)
	}

	// Almacenar en caché
	asistenciaEventoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Lista de asistencias obtenida con éxito", data)
}

func GetAsistenciaEventoID(c *fiber.Ctx) error {
	// Verificar caché primero
	if result, exists := asistenciaEventoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Asistencia obtenida del caché", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.GetID(c, models.GetAsistenciaEventoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener la asistencia", err)
	}

	// Almacenar en caché
	asistenciaEventoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Asistencia obtenida con éxito", data)
}

func PostAsistenciaEvento(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de asistencia no válidos", err)
	}

	// Seguridad
	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para registrar asistencia", err)
	}

	// Procesar datos
	err = handlers.HandlePost(c, models.PostAsistenciaEvento)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar asistencia", err)
	}

	// Invalidar caché
	asistenciaEventoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Asistencia registrada con éxito", nil)
}

func PutAsistenciaEvento(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de asistencia no válidos", err)
	}

	// Seguridad
	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para actualizar asistencia", err)
	}

	// Procesar datos
	err = handlers.Put(c, models.PutAsistenciaEvento)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar asistencia", err)
	}

	// Invalidar caché
	asistenciaEventoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Asistencia actualizada con éxito", nil)
}

func DeleteAsistenciaEvento(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteAsistenciaEvento)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar asistencia", err)
	}

	// Invalidar caché
	asistenciaEventoSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Asistencia eliminada con éxito", nil)
}