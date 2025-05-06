package interfaz

import (
	"claseobrera/security/object"
	"claseobrera/security/operation"
	"claseobrera/security/utils"
)

// Authorize es la función principal que se encarga de autorizar y validar los datos entrantes.
func AuthorizeData(clientIP string, data interface{}) (bool, error) {
	parsedData, err := object.ParseData(data)
	if err != nil {
		return false, err
	}
	return operation.ValidateData(clientIP, parsedData)
}

// GenerateNonce genera un nonce único para proteger contra replay attacks.
func GenerateNonce(clientIP string) string {
	return utils.GenerateNonce(clientIP)
}
