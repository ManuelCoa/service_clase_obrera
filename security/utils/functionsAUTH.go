package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

// Base64UrlDecode decodifica una cadena Base64 URL.
// Agrega padding en caso de ser necesario.
func Base64UrlDecode(seg string) ([]byte, error) {
	// Agregar padding si es necesario
	switch len(seg) % 4 {
	case 2:
		seg += "=="
	case 3:
		seg += "="
	}
	return base64.URLEncoding.DecodeString(seg)
}

// ComputeHMACSHA256 calcula el HMAC SHA256 de un mensaje usando una clave secreta.
func ComputeHMACSHA256(message string, secret []byte) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	// Codificar en Base64 URL sin padding
	encoded := base64.URLEncoding.EncodeToString(expectedMAC)
	return strings.TrimRight(encoded, "=")
}

// TrimBearerToken elimina el prefijo "Bearer " de un token, si existe.
func TrimBearerToken(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
