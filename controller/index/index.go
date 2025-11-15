package index

import (
	"net/http"

	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
)

func Index(w http.ResponseWriter, r *http.Request) {
	p := viewbackend.FrontEndParams{
		Title:   "Base",
		Message: "This is a new beginning! Hello from Index",
	}
	viewbackend.Frontend_Home(w, p)
}
