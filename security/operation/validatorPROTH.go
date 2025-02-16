package operation

import (
	"errors"

	"claseobrera/security/utils"
)

func ValidateData(input string) (string, error) {

	if !utils.ValidateUTF8(input) {
		return "", errors.New("datos binarios inconsistentes: la cadena no es UTF-8 válida")
	}

	if utils.ContainsSQLInjection(input) {
		sanitized := utils.SanitizeInput(input)
		if utils.ContainsSQLInjection(sanitized) {
			return "", errors.New("la entrada contiene patrones potencialmente peligrosos de inyección SQL")
		}
		return sanitized, nil
	}

	return input, nil
}
