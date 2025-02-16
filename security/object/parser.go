package object

import (
	"encoding/json"
	"errors"
)

func ParseInput(input interface{}) (string, error) {
	switch v := input.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return "", errors.New("no se pudo parsear la entrada")
		}
		return string(b), nil
	}
}
