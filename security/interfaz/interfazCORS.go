package interfaz

import (
	"claseobrera/security/object"
	"claseobrera/security/operation"
)

// CORSConfig se reexporta para que el consumidor use la configuración a través de la interfaz,
// sin acceder directamente al package interno object.
type CORSConfig = object.CORSConfig

// BuildCORSHeaders es la función pública que, dada una configuración CORS, un origin y un método HTTP,
// devuelve el mapa de headers que se deben incluir en la respuesta y un flag que indica si es una solicitud preflight.
func BuildCORSHeaders(config CORSConfig, origin, method string) (map[string]string, bool) {
	return operation.ProcessCORS(config, origin, method)
}
