package object

import (
	"encoding/json"
	"errors"
	"strings"

	"claseobrera/security/utils" // Ajusta el path de acuerdo a tu módulo
)

// JWTToken representa la estructura de un JWT.
type JWTToken struct {
	Raw          string
	Header       map[string]interface{}
	Payload      map[string]interface{}
	Signature    string
	SigningInput string // Cadena "header.payload" para la verificación de la firma.
}

// ParseToken convierte el string de un JWT en un objeto JWTToken.
func ParseToken(tokenString string) (*JWTToken, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("token malformado")
	}

	headerSegment, payloadSegment, signatureSegment := parts[0], parts[1], parts[2]

	headerBytes, err := utils.Base64UrlDecode(headerSegment)
	if err != nil {
		return nil, errors.New("error decodificando header")
	}

	payloadBytes, err := utils.Base64UrlDecode(payloadSegment)
	if err != nil {
		return nil, errors.New("error decodificando payload")
	}

	var header map[string]interface{}
	var payload map[string]interface{}

	if err = json.Unmarshal(headerBytes, &header); err != nil {
		return nil, errors.New("error parseando header JSON")
	}

	if err = json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, errors.New("error parseando payload JSON")
	}

	token := &JWTToken{
		Raw:          tokenString,
		Header:       header,
		Payload:      payload,
		Signature:    signatureSegment,
		SigningInput: headerSegment + "." + payloadSegment,
	}

	return token, nil
}
