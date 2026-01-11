package template

import (
	"html/template"
	"strings"
)

var Funcs = template.FuncMap{
	"uppercase": ToUpper,
	"truncate":  Truncate,
	"add":       Add,
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

func Add(first int, second int) int {
	return first + second
}
