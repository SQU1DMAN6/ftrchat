package index

import (
	"net/http"

	"fmt"

	"github.com/SQU1DMAN6/ftrchat/config"

	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
)

func Index(w http.ResponseWriter, r *http.Request) {
	paramData := viewbackend.FrontEndParams{
		Title: "Base",
	}

	SS := config.GetSessionManager()

	msg := SS.GetString(r.Context(), "message")
	isLoggedIn := SS.GetBool(r.Context(), "isLoggedIn")
	userEmail := SS.GetString(r.Context(), "userEmail")
	userName := SS.GetString(r.Context(), "name")

	fmt.Println("Message (From Index): ", msg)
	fmt.Println("userEmail (From Index): ", userEmail)
	fmt.Println("User Name (From Index): ", userName)

	paramData.Authenticated = isLoggedIn
	paramData.Message = fmt.Sprintf("Welcome, %s, to the FtR Project.", userName)

	if isLoggedIn != true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	viewbackend.Frontend_Home(w, paramData)
}
