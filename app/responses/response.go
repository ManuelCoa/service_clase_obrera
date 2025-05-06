package responses

import "github.com/gofiber/fiber/v2"

/*
	ErrorDetail representa la estructura para los detalles de un error.
	Incluye un código de error y un mensaje descriptivo.
*/

type ErrorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/*
	Meta representa la estructura de meta información para una respuesta.
	Incluye el ID de la solicitud (opcional) y la versión de la API (opcional).
*/

type Meta struct {
	RequestID  interface{} `json:"request_id,omitempty"`
	APIVersion string      `json:"api_version,omitempty"`
}

/*
	CreateResponse genera una respuesta JSON estándar para las solicitudes HTTP.
	Crea un mapa con la información de estado, código de respuesta, mensaje, datos,
	detalles del error (si los hay), y meta información. Luego envía la respuesta
	con el código de estado especificado.
*/

func CreateResponse(c *fiber.Ctx, status string, code int, message string, data interface{}, errDetail *ErrorDetail, meta *Meta) error {
	response := fiber.Map{
		"status":  status,
		"code":    code,
		"message": message,
		"data":    data,
		"error":   errDetail,
		"meta":    meta,
	}

	return c.Status(code).JSON(response)
}
