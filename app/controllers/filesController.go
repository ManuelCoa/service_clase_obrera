package controllers

import (
	observer "claseobrera/app/cache"
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	caching "claseobrera/caching/interfaz"
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

var filesCache = caching.NewCache("Files", 100)
var filesSubject = observer.NewSubject()

func init() {
	filesSubject.AddObserver(updateFilesCache)
}

func updateFilesCache(url string) {
	filesCache.Delete(url)
}

func GetFiles(c *fiber.Ctx) error {
	if result, exists := filesCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Archivos obtenidos del cache", result)
	}

	data, err := handlers.Get(c, models.GetFiles)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener archivos", err)
	}

	filesCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Archivos obtenidos exitosamente", data)
}

func GetFileID(c *fiber.Ctx) error {
	if result, exists := filesCache.Get(c.OriginalURL()); exists {
		return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Archivo obtenido del cache", result)
	}

	data, err := handlers.GetStringID(c, models.GetFileID)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al obtener archivo", err)
	}

	filesCache.Set(c.OriginalURL(), data, len(c.OriginalURL()))
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Archivo obtenido exitosamente", data)
}

func PostFile(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de archivo no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para subir archivos", err)
	}

	err = handlers.HandlePost(c, models.PostFile)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al subir archivo", err)
	}

	filesSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusCreated, "Archivo subido exitosamente", nil)
}

func PutFile(c *fiber.Ctx) error {
	clientIP := c.IP()
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusBadRequest, "Datos de archivo no válidos", err)
	}

	data["nonce"] = interfaz.GenerateNonce(clientIP)
	authorized, err := interfaz.AuthorizeData(clientIP, data)
	if !authorized {
		return handlers.CreateErrorResponse(c, fiber.StatusUnauthorized, "No autorizado para modificar archivos", err)
	}

	err = handlers.Put(c, models.PutFile)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al modificar archivo", err)
	}

	filesSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Archivo modificado exitosamente", nil)
}

func DeleteFile(c *fiber.Ctx) error {
	err := handlers.DeleteString(c, models.DeleteFile)
	if err != nil {
		return handlers.CreateErrorResponse(c, fiber.StatusInternalServerError, "Error al eliminar archivo", err)
	}

	filesSubject.NotifyObservers(c.OriginalURL())
	return handlers.CreateSuccessResponse(c, fiber.StatusOK, "Archivo eliminado exitosamente", nil)
}