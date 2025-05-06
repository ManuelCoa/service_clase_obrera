package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var viviendaCache = caching.NewCache("Vivienda", 100)
var viviendaSubject = observer.NewSubject()

func init() {
	viviendaSubject.AddObserver(updateViviendaCache)
}

// updateViviendaCache es una función observadora que actualiza la caché
func updateViviendaCache(url string) {
	viviendaCache.Delete(url)
}

func GetVivienda(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := viviendaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos de vivienda obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.Get(c, models.GetVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener viviendas", err)
	}

	// Almacenar los datos en la caché
	viviendaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Listado de viviendas obtenido con éxito", data)
}

func GetViviendaID(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := viviendaCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos de vivienda obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.GetID(c, models.GetViviendaID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener la vivienda", err)
	}

	// Almacenar los datos en la caché
	viviendaCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Vivienda obtenida con éxito", data)
}

func PostVivienda(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	// Analizar el cuerpo de la solicitud
	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de vivienda no válidos", err)
	}

	// Generar nonce y autorizar datos
	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para registrar vivienda", err)
	}

	// Procesar los datos
	err = handlers.HandlePost(c, models.PostVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar la vivienda", err)
	}

	// Notificar a los observadores para actualizar la caché
	viviendaSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Vivienda registrada con éxito", nil)
}

func PutVivienda(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	// Analizar el cuerpo de la solicitud
	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de vivienda no válidos", err)
	}

	// Generar nonce y autorizar datos
	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para actualizar vivienda", err)
	}

	// Procesar los datos
	err = handlers.Put(c, models.PutVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar la vivienda", err)
	}

	// Notificar a los observadores para actualizar la caché
	viviendaSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Vivienda actualizada con éxito", nil)
}

func DeleteVivienda(c *fiber.Ctx) error {
	// Procesar los datos
	err := handlers.Delete(c, models.DeleteVivienda)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar la vivienda", err)
	}

	// Notificar a los observadores para actualizar la caché
	viviendaSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Vivienda eliminada con éxito", nil)
}