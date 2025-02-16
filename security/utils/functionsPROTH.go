package utils

import (
	"strings"
	"unicode/utf8"
)

func ValidateUTF8(input string) bool {
	return utf8.ValidString(input)
}

func ContainsSQLInjection(input string) bool {

	patrones := []string{
		"DROP", "SELECT", "INSERT", "DELETE", "UPDATE", "--", ";", "/*", "*/", "' OR", "\" OR",
	}
	inputUpper := strings.ToUpper(input)
	for _, p := range patrones {
		if strings.Contains(inputUpper, p) {
			return true
		}
	}
	return false
}

func SanitizeInput(input string) string {
	peligrosos := []string{";", "--", "/*", "*/", "' OR", "\" OR"}
	resultado := input
	for _, p := range peligrosos {
		resultado = strings.ReplaceAll(resultado, p, "")
	}
	return resultado
}
