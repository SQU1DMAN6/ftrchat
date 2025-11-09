package index

import (
	"net/http"

	viewbackend "github.com/SQU1DMAN6/ftrchat/html/view/backend"
)

func Other(w http.ResponseWriter, r *http.Request) {
	p := viewbackend.FrontEndParams{
		Title:   "Base",
		Message: "This is a new beginning! Hello from Other",
	}
	viewbackend.Frontend_Other(w, p)
}
