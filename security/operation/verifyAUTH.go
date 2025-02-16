package operation

import (
	"errors"

	"claseobrera/security/object"
	"claseobrera/security/utils"
)

// VerifyToken verifica la firma de un token JWT utilizando el secreto proporcionado.
func VerifyToken(token *object.JWTToken, secret []byte) (bool, error) {
	computedSignature := utils.ComputeHMACSHA256(token.SigningInput, secret)
	if computedSignature != token.Signature {
		return false, errors.New("firma inválida")
	}
	return true, nil
}
