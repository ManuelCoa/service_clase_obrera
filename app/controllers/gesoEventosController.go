package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var geoEventoCache = caching.NewCache("Knapsack", 100)
var geoEventoSubject = observer.NewSubject()

func init() {
	geoEventoSubject.AddObserver(updateGeoEventoCache)
}

// updateGeoEventoCache es una función observadora que actualiza la caché
func updateGeoEventoCache(url string) {
	geoEventoCache.Delete(url)
}

func GetGeoEvento(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := geoEventoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos geo-eventos obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.Get(c, models.GetGeoEvento)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener geo-eventos", err)
	}

	// Almacenar los datos en la caché
	geoEventoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos geo-eventos obtenidos con éxito", data)
}

func GetGeoEventoID(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := geoEventoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos geo-evento obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.GetID(c, models.GetGeoEventoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener geo-evento", err)
	}

	// Almacenar los datos en la caché
	geoEventoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos geo-evento obtenidos con éxito", data)
}

func PostGeoEvento(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	// Analizar el cuerpo de la solicitud
	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de geo-evento no válidos", err)
	}

	// Generar nonce y autorizar datos
	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para crear geo-evento", err)
	}

	// Procesar los datos
	err = handlers.HandlePost(c, models.PostGeoEvento)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al crear geo-evento", err)
	}

	// Notificar a los observadores para actualizar la caché
	geoEventoSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Geo-evento creado con éxito", nil)
}

func PutGeoEvento(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	// Analizar el cuerpo de la solicitud
	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de geo-evento no válidos", err)
	}

	// Generar nonce y autorizar datos
	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado para actualizar geo-evento", err)
	}

	// Procesar los datos
	err = handlers.Put(c, models.PutGeoEvento)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar geo-evento", err)
	}

	// Notificar a los observadores para actualizar la caché
	geoEventoSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Geo-evento actualizado con éxito", nil)
}

func DeleteGeoEvento(c *fiber.Ctx) error {
	// Procesar los datos
	err := handlers.Delete(c, models.DeleteGeoEvento)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar geo-evento", err)
	}

	// Notificar a los observadores para actualizar la caché
	geoEventoSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Geo-evento eliminado con éxito", nil)
}