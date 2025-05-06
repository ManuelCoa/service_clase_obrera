package operation

import (
	"claseobrera/security/object"
	"claseobrera/security/utils"
	"strconv"
)

// ProcessCORS procesa la solicitud CORS según la configuración recibida, el origin y el método HTTP.
// Devuelve un mapa de headers que deben incluirse en la respuesta y un booleano que indica si se trata de una
// solicitud preflight (método OPTIONS).
func ProcessCORS(config object.CORSConfig, origin, method string) (map[string]string, bool) {
	headers := make(map[string]string)

	// Si no se provee origin o éste no está permitido, no se añaden headers CORS.
	if origin == "" || !utils.IsOriginAllowed(origin, config.AllowedOrigins) {
		return headers, false
	}

	// Si se permite cualquier origin ("*") y no se requieren credenciales, se asigna "*".
	if len(config.AllowedOrigins) == 1 && config.AllowedOrigins[0] == "*" && !config.AllowCredentials {
		headers["Access-Control-Allow-Origin"] = "*"
	} else {
		// En otro caso se responde con el origin de la solicitud.
		headers["Access-Control-Allow-Origin"] = origin
	}

	if config.AllowCredentials {
		headers["Access-Control-Allow-Credentials"] = "true"
	}

	// Si es una solicitud preflight, se añaden headers específicos.
	if method == "OPTIONS" {
		headers["Access-Control-Allow-Methods"] = utils.JoinStrings(config.AllowedMethods)
		headers["Access-Control-Allow-Headers"] = utils.JoinStrings(config.AllowedHeaders)
		if config.MaxAge > 0 {
			headers["Access-Control-Max-Age"] = strconv.Itoa(config.MaxAge)
		}
		return headers, true
	}

	// Para solicitudes simples, se pueden exponer headers adicionales.
	if len(config.ExposedHeaders) > 0 {
		headers["Access-Control-Expose-Headers"] = utils.JoinStrings(config.ExposedHeaders)
	}

	return headers, false
}
