package interfaz

import (
	"claseobrera/security/object"
	"claseobrera/security/operation"
	"claseobrera/security/utils"
)

// Authorize es la función de interfaz que valida un JWT y retorna sus claims (payload).
func Authorize(tokenString string, secret []byte) (map[string]interface{}, error) {
	// Eliminar prefijo "Bearer " si existe.
	tokenString = utils.TrimBearerToken(tokenString)

	// Parsear el token.
	token, err := object.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	// Verificar la firma.
	valid, err := operation.VerifyToken(token, secret)
	if !valid {
		return nil, err
	}

	// Retornar el payload (claims) del token.
	return token.Payload, nil
}
