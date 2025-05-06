package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var amonestacionesCache = caching.NewCache("Amonestaciones", 100)
var amonestacionesSubject = observer.NewSubject()

func init() {
	amonestacionesSubject.AddObserver(updateAmonestacionesCache)
}

// updateAmonestacionesCache es una función observadora que actualiza la caché
func updateAmonestacionesCache(url string) {
	amonestacionesCache.Delete(url)
}

func GetAmonestaciones(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := amonestacionesCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.Get(c, models.GetAmonestaciones)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Almacenar los datos en la caché
	amonestacionesCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func GetAmonestacionID(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := amonestacionesCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.GetID(c, models.GetAmonestacionID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Almacenar los datos en la caché
	amonestacionesCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func PostAmonestacion(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	// Analizar el cuerpo de la solicitud
	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de entrada no válidos", err)
	}

	// Generar nonce y autorizar datos
	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	// Procesar los datos
	err = handlers.HandlePost(c, models.PostAmonestacion)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Notificar a los observadores para actualizar la caché
	amonestacionesSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "La amonestación se ha creado con éxito", nil)
}

func PutAmonestacion(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	// Analizar el cuerpo de la solicitud
	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de entrada no válidos", err)
	}

	// Generar nonce y autorizar datos
	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	// Procesar los datos
	err = handlers.Put(c, models.PutAmonestacion)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Notificar a los observadores para actualizar la caché
	amonestacionesSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "La amonestación se ha actualizado con éxito", nil)
}

func DeleteAmonestacion(c *fiber.Ctx) error {
	// Procesar los datos
	err := handlers.Delete(c, models.DeleteAmonestacion)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Notificar a los observadores para actualizar la caché
	amonestacionesSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "La amonestación se ha eliminado con éxito", nil)
}