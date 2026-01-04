package template

import (
	"html/template"
	"strings"
)

var Funcs = template.FuncMap{
	"uppercase": ToUpper,
	"truncate":  Truncate,
}

func ToUpper(v string) string {
	return strings.ToUpper(v)
}

func Truncate(stringToTruncate string) string {
	runes := []rune(stringToTruncate)
	if len(runes) > 500 {
		return string(runes[:500]) + "..."
	}

	return stringToTruncate
}
