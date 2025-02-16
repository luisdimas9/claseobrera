package interfaz

import (
	"claseobrera/security/object"
	"claseobrera/security/operation"
)

func AuthorizeData(input interface{}) (string, error) {
	parsed, err := object.ParseInput(input)
	if err != nil {
		return "", err
	}
	sanitized, err := operation.ValidateData(parsed)
	if err != nil {
		return "", err
	}
	return sanitized, nil
}
