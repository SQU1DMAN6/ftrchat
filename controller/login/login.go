package login

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SQU1DMAN6/ftrchat/config"
	model_user "github.com/SQU1DMAN6/ftrchat/model"
	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
)

func LoginMain(w http.ResponseWriter, r *http.Request) {
	p := viewbackend.FrontEndParams{
		Title:   "Login",
		Message: "Log in to an existing FtR account",
		Error:   make(map[string]string),
	}

	viewbackend.LoginMain(w, p)
}

func LoginMainPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "ParseForm() error", http.StatusBadRequest)
		return
	}

	userEmail := strings.TrimSpace(r.FormValue("auth_email"))
	userPassword := strings.TrimSpace(r.FormValue("auth_password"))

	if userEmail == "" || userPassword == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	fmt.Println("User email in form:", userEmail)
	fmt.Println("User password in form:", userPassword)

	db := config.GetDB()

	user, err := model_user.CheckPassword(db, userEmail, userPassword)
	if err != nil {
		fmt.Println("Error:", err)
		paramData := viewbackend.FrontEndParams{
			Title:   "Login",
			Message: "Log in to an existing FtR account",
			Error:   make(map[string]string),
		}

		paramData.Error["general"] = "Something wrong, your turn your turn your turn your turn your turn..."
		viewbackend.LoginMain(w, paramData)
		return

	}

	fmt.Println(user)

	SS := config.GetSessionManager()	

	SS.Put(r.Context(), "message", "Hello from a session!")
	SS.Put(r.Context(), "userEmail", userEmail)
	SS.Put(r.Context(), "isLogged", true)

	// Setting the session with the user data
	// Redirect

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
