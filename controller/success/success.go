package success

import (
	"net/http"

	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
)

func SuccessRegister(w http.ResponseWriter, r *http.Request) {
	p := viewbackend.FrontEndParams{
		Title:   "Success",
		Message: "Successfully registered for a new account",
	}

	viewbackend.SuccessRegister(w, p)
}
