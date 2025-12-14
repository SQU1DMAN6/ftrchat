package viewbackend

import "html/template"

type FrontEndParams struct {
	Title         string
	Name          string
	Message       string
	UserData      interface{}
	SessionData   map[string]string
	CurrentURL    string
	Page          int
	CSRFToken     template.HTML
	Pagination    map[string]interface{}
	Authenticated bool
	Error         map[string]string
}
