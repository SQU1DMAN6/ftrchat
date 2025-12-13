package index

import (
	"net/http"

	"github.com/SQU1DMAN6/ftrchat/config"
	"fmt"

	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
)

func Index(w http.ResponseWriter, r *http.Request) {
	paramData := viewbackend.FrontEndParams{
		Title:   "Base",
		Message: "This is a new beginning! Hello from Index",
		
	}

	
	SS := config.GetSessionManager()	

	msg := SS.GetString(r.Context(), "message")
	isLogged := SS.GetBool(r.Context(), "isLogged")
	userEmail := SS.GetString(r.Context(), "userEmail")
	// io.WriteString(w, msg)

	fmt.Println("msg", msg)
	fmt.Println("userEmail", userEmail)

	paramData.Authenticated = isLogged
	

	viewbackend.Frontend_Home(w, paramData)
}
