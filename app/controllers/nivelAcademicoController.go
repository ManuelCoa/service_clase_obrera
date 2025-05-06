package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var nivelAcademicoCache = caching.NewCache("Knapsack", 100)
var nivelAcademicoSubject = observer.NewSubject()

func init() {
	nivelAcademicoSubject.AddObserver(updateNivelAcademicoCache)
}

// updateCache es una función observadora que actualiza la caché
func updateNivelAcademicoCache(url string) {
	nivelAcademicoCache.Delete(url)
}

func GetNivelAcademico(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := nivelAcademicoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.Get(c, models.GetNivelAcademico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Almacenar los datos en la caché
	nivelAcademicoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func GetNivelAcademicoID(c *fiber.Ctx) error {
	// Verificar si los datos están en la caché
	if result, exists := nivelAcademicoCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	// Obtener datos de la base de datos
	data, err := handlers.GetID(c, models.GetNivelAcademicoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Almacenar los datos en la caché
	nivelAcademicoCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	// Devolver los datos al cliente
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func PostNivelAcademico(c *fiber.Ctx) error {
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
	err = handlers.HandlePost(c, models.PostNivelAcademico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Notificar a los observadores para actualizar la caché
	nivelAcademicoSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "El nivel académico se ha creado con éxito", nil)
}

func PutNivelAcademico(c *fiber.Ctx) error {
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
	err = handlers.Put(c, models.PutNivelAcademico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Notificar a los observadores para actualizar la caché
	nivelAcademicoSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "El nivel académico se ha actualizado con éxito", nil)
}

func DeleteNivelAcademico(c *fiber.Ctx) error {
	// Procesar los datos
	err := handlers.Delete(c, models.DeleteNivelAcademico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	// Notificar a los observadores para actualizar la caché
	nivelAcademicoSubject.NotifyObservers(c.OriginalURL())

	// Devolver una respuesta exitosa
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "El nivel académico se ha eliminado con éxito", nil)
}