package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var asistenciaCache = caching.NewCache("Asistencia", 100)
var asistenciaSubject = observer.NewSubject()

func init() {
	asistenciaSubject.AddObserver(updateAsistenciaCache)
}

// updateAsistenciaCache es una función observadora que actualiza la caché
func updateAsistenciaCache(url string) {
	asistenciaCache.Delete(url)
}

func GetAsistencias(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := asistenciaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos de asistencia obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.Get(c, models.GetAsistencias)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener los registros de asistencia", err)
	}

	// Almacenar los datos en la caché
	asistenciaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registros de asistencia obtenidos con éxito", data)
}

func GetAsistenciaID(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := asistenciaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos de asistencia obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.GetID(c, models.GetAsistenciaID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener el registro de asistencia", err)
	}

	// Almacenar los datos en la caché
	asistenciaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Registro de asistencia obtenido con éxito", data)
}

func PostAsistencia(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	// Analizar el cuerpo de la solicitud
	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de asistencia no válidos", err)
	}

	// Generar nonce y autorizar datos
	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para registrar asistencia", err)
	}

	// Procesar los datos
	err = handlers.HandlePost(c, models.PostAsistencia)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar la asistencia", err)
	}

	// Notificar a los observadores para actualizar la caché
	asistenciaSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Asistencia registrada con éxito", nil)
}

func PutAsistencia(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	// Analizar el cuerpo de la solicitud
	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de asistencia no válidos", err)
	}

	// Generar nonce y autorizar datos
	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para actualizar asistencia", err)
	}

	// Procesar los datos
	err = handlers.Put(c, models.PutAsistencia)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar el registro de asistencia", err)
	}

	// Notificar a los observadores para actualizar la caché
	asistenciaSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Asistencia actualizada con éxito", nil)
}

func DeleteAsistencia(c *fiber.Ctx) error {
	// Procesar los datos
	err := handlers.Delete(c, models.DeleteAsistencia)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar el registro de asistencia", err)
	}

	// Notificar a los observadores para actualizar la caché
	asistenciaSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Asistencia eliminada con éxito", nil)
}