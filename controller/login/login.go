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

		paramData.Error["general"] = fmt.Sprintf("Encountered an error trying to log in: %s", err)
		viewbackend.LoginMain(w, paramData)
		return
	}

	fmt.Printf("\nUser: %v | ID: %v | Email: %v\n", user.Name, user.ID, user.Email)

	SS := config.GetSessionManager()

	SS.Put(r.Context(), "message", "Hello, fellow people with ding-a-lings!")
	SS.Put(r.Context(), "userEmail", user.Email)
	SS.Put(r.Context(), "name", user.Name)
	SS.Put(r.Context(), "isLoggedIn", true)

	// Setting the session with the user data
	// Redirect

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
