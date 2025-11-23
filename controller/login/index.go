package login

import (
	"fmt"
	"net/http"
	"strings"

	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
)

func LoginMain(w http.ResponseWriter, r *http.Request) {
	p := viewbackend.FrontEndParams{
		Title:   "Login",
		Message: "Log in to an existing FtR account",
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
	fmt.Println("User email:", userEmail)
	fmt.Println("User password:", userPassword)

	// TODO: Implement user authentication logic
}
