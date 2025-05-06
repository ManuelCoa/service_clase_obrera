package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var organizacionGeneralCache = caching.NewCache("Knapsack", 100)
var organizacionGeneralSubject = observer.NewSubject()

func init() {
	organizacionGeneralSubject.AddObserver(updateOrganizacionGeneralCache)
}

// updateCache es una función observadora que actualiza la caché
func updateOrganizacionGeneralCache(url string) {
	organizacionGeneralCache.Delete(url)
}

func GetOrganizacionGeneral(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := organizacionGeneralCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.Get(c, models.GetOrganizacionGeneral)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Almacenar los datos en la caché
	organizacionGeneralCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func GetOrganizacionGeneralID(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := organizacionGeneralCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.GetID(c, models.GetOrganizacionGeneralID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Almacenar los datos en la caché
	organizacionGeneralCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func PostOrganizacionGeneral(c *fiber.Ctx) error {
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
	err = handlers.HandlePost(c, models.PostOrganizacionGeneral)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Notificar a los observadores para actualizar la caché
	organizacionGeneralSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "La organización general se ha creado con éxito", nil)
}

func PutOrganizacionGeneral(c *fiber.Ctx) error {
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
	err = handlers.Put(c, models.PutOrganizacionGeneral)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Notificar a los observadores para actualizar la caché
	organizacionGeneralSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "La organización general se ha actualizado con éxito", nil)
}

func DeleteOrganizacionGeneral(c *fiber.Ctx) error {
	// Procesar los datos
	err := handlers.Delete(c, models.DeleteOrganizacionGeneralID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Notificar a los observadores para actualizar la caché
	organizacionGeneralSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "La organización general se ha eliminado con éxito", nil)
}