package template

import (
	"embed"
	"html/template"
	"sync"
)

//go:embed themes/backend/*/*.html
var filesystem embed.FS

var (
	cachedTemplate *template.Template
	once           sync.Once
	devMode        = true // Set to true during development
)

func ParseBackEnd(files ...string) *template.Template {
	allFiles := append(
		[]string{
			"themes/backend/layout/base.html", //==> the template will have the name "base_pure.html"

		},
		files...)
	return template.Must(
		template.New("").Funcs(Funcs).ParseFS(filesystem, allFiles...))
}

func ParseBackEndLogin(files ...string) *template.Template {
	allFiles := append(
		[]string{
			"themes/backend/layout/baselogin.html", //==> the template will have the name "base_pure.html"

		},
		files...)
	return template.Must(
		template.New("").Funcs(Funcs).ParseFS(filesystem, allFiles...))
}
