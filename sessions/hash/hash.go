package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

// CheckPassword compara el hash de una contraseña con el hash almacenado
func CheckPassword(providedPassword, storedPassword string) bool {
	hashedPassword, err := HashPassword(providedPassword)
	if err != nil {
		return false
	}
	return hashedPassword == storedPassword
}
