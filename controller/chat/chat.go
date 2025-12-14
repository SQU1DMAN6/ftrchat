package chat

import (
	"fmt"
	"net/http"

	"github.com/SQU1DMAN6/ftrchat/config"
	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
)

func ChatMain(w http.ResponseWriter, r *http.Request) {
	paramData := viewbackend.FrontEndParams{
		Title:   "Chat",
		Message: "This is where all the chat thingies go.",
	}

	SS := config.GetSessionManager()

	msg := SS.GetString(r.Context(), "message")
	isLoggedIn := SS.GetBool(r.Context(), "isLogged")
	userEmail := SS.GetString(r.Context(), "userEmail")

	fmt.Println("msg", msg)
	fmt.Println("userEmail", userEmail)

	if isLoggedIn != true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	viewbackend.Frontend_ChatMain(w, paramData)
}
