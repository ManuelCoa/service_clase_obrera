package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var transportePublicocache = caching.NewCache("Knapsack", 100)
var transportePublicoSubject = observer.NewSubject()

func init() {
	transportePublicoSubject.AddObserver(updateTransportePublicoCache)
}

func updateTransportePublicoCache(url string) {
	transportePublicocache.Delete(url)
}

func GetTransportePublico(c *fiber.Ctx) error {
	if result, exists := transportePublicocache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	data, err := handlers.Get(c, models.GetTransportePublico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	transportePublicocache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func GetTransportePublicoID(c *fiber.Ctx) error {
	if result, exists := transportePublicocache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito del cache", result)
	}

	data, err := handlers.GetID(c, models.GetTransportePublicoID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	transportePublicocache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func PostTransportePublico(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de entrada no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.HandlePost(c, models.PostTransportePublico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	transportePublicoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "El transporte público se ha creado con éxito", nil)
}

func PutTransportePublico(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de entrada no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "Acceso no autorizado", err)
	}

	err = handlers.Put(c, models.PutTransportePublico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	transportePublicoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "El transporte público se ha actualizado con éxito", nil)
}

func DeleteTransportePublico(c *fiber.Ctx) error {
	err := handlers.Delete(c, models.DeleteTransportePublico)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}

	transportePublicoSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "El transporte público se ha eliminado con éxito", nil)
}