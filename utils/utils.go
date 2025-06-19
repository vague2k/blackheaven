package utils

import (
	"fmt"
	"unicode"
)

func ID(formID, name, component, element string) string {
	return fmt.Sprintf("%s-%s-%s-%s", formID, name, component, element)
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
