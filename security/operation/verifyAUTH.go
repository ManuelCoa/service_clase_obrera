package operation

import (
	"errors"

	"claseobrera/security/object"
	"claseobrera/security/utils"
)

func VerifyToken(token *object.JWTToken, secret []byte) (bool, error) {
	computedSignature := utils.ComputeHMACSHA256(token.SigningInput, secret)

	if computedSignature != token.Signature {
		return false, errors.New("firma inválida")
	}
	return true, nil
}
