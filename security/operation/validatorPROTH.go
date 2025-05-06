package operation

import (
	"claseobrera/security/utils"
	"errors"
	"sync"
	"time"
)

var (
	rateLimiter  = utils.NewRateLimiter(100, time.Minute)
	storedNonces = make(map[string]bool)
	nonceMu      sync.Mutex
)

// ValidateData valida los datos parseados.
func ValidateData(clientIP string, data map[string]interface{}) (bool, error) {
	if !rateLimiter.Allow(clientIP) {
		return false, errors.New("límite de solicitudes excedido")
	}

	nonce, ok := data["nonce"].(string)
	if !ok || !utils.ValidateNonce(nonce, storedNonces) {
		return false, errors.New("nonce inválido o reutilizado")
	}

	if utils.ContainsSQLInjection(data) {
		return false, errors.New("posible inyección SQL detectada")
	}
	if utils.ContainsNoSQLInjection(data) {
		return false, errors.New("posible inyección NoSQL detectada")
	}
	if utils.ContainsCommandInjection(data) {
		return false, errors.New("posible inyección de comandos detectada")
	}
	if utils.ContainsXMLInjection(data) {
		return false, errors.New("posible inyección XML detectada")
	}
	if utils.ContainsJSONInjection(data) {
		return false, errors.New("posible inyección JSON detectada")
	}
	if utils.ContainsBinaryData(data) {
		return false, errors.New("datos binarios inconsistentes detectados")
	}
	if utils.ContainsDamagedData(data) {
		return false, errors.New("datos dañados detectados")
	}
	return true, nil
}
