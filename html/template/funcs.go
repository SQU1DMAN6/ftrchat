package template

import (
	"html/template"
	"strings"
)

var Funcs = template.FuncMap{
	"uppercase": func(v string) string {
		return strings.ToUpper(v)
	},
}
