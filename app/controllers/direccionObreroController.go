package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var direccionObreroCache = caching.NewCache("DireccionObrero", 100)
var direccionObreroSubject = observer.NewSubject()

func init() {
	direccionObreroSubject.AddObserver(updateDireccionObreroCache)
}

func updateDireccionObreroCache(url string) {
	direccionObreroCache.Delete(url)
}

func GetDireccionesObrero(c *fiber.Ctx) error {
	if result, exists := direccionObreroCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Direcciones obtenidas del cache", result)
	}

	data, err := handlers.Get(c, models.GetDireccionesObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener direcciones", err)
	}

	direccionObreroCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Direcciones obtenidas con éxito", data)
}

func GetDireccionObreroID(c *fiber.Ctx) error {
	if result, exists := direccionObreroCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Dirección obtenida del cache", result)
	}

	data, err := handlers.GetID(c, models.GetDireccionObreroID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener dirección", err)
	}

	direccionObreroCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Dirección obtenida con éxito", data)
}
//funcion para consultar por id de institucion 
func GetDireccionObreroIDInstitucion(c *fiber.Ctx) error {

	if result, exists := obreroCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos del obrero obtenidos del caché", result)
	}

	data, err := handlers.GetIDInstitucion(c, models.GetDireccionObreroInstitucionID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener datos del obrero", err)
	}

	obreroCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos del obrero obtenidos exitosamente", data)
}

func PostDireccionObrero(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de dirección inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostDireccionObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al registrar dirección", err)
	}

	direccionObreroSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Dirección registrada exitosamente", nil)
}

func PutDireccionObrero(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de dirección inválidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutDireccionObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al actualizar dirección", err)
	}

	direccionObreroSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Dirección actualizada exitosamente", nil)
}

func DeleteDireccionObrero(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteDireccionObrero)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar dirección", err)
	}

	direccionObreroSubject.NotifyObservers(c.OriginalURL())

	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Dirección eliminada exitosamente", nil)
}