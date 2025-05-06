package middleware

import (
	"claseobrera/security/interfaz"

	"github.com/gofiber/fiber/v2"
)

// convertCorsOptions convierte las opciones locales a la configuración esperada por el módulo security.
func convertCorsOptions(options CorsOptions) interfaz.CORSConfig {
	return interfaz.CORSConfig{
		AllowedOrigins:   options.AllowedOrigins,
		AllowedMethods:   options.AllowedMethods,
		AllowedHeaders:   options.AllowedHeaders,
		AllowCredentials: options.AllowCredentials,
		ExposedHeaders:   options.ExposedHeaders,
		MaxAge:           options.MaxAge,
	}
}

// CorsMiddleware retorna un handler de Fiber que aplica la configuración CORS usando el módulo security.
// El router utilizará este middleware pasando un CorsOptions, sin conocer nada del módulo security.
func CorsMiddleware(options CorsOptions) fiber.Handler {
	// Se convierte la configuración local al tipo requerido por el módulo security.
	config := convertCorsOptions(options)

	return func(c *fiber.Ctx) error {
		origin := c.Get("Origin")
		method := c.Method()

		// Se utiliza la interfaz de security para obtener los headers CORS.
		headers, isPreflight := interfaz.BuildCORSHeaders(config, origin, method)

		// Se asignan los headers a la respuesta.
		for key, value := range headers {
			c.Set(key, value)
		}

		// Si es una solicitud preflight, se responde sin continuar.
		if isPreflight {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	}
}
