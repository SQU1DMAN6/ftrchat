package register

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SQU1DMAN6/ftrchat/config"
	model_user "github.com/SQU1DMAN6/ftrchat/model"
	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
)

func RegisterMain(w http.ResponseWriter, r *http.Request) {
	p := viewbackend.FrontEndParams{
		Title:   "Register",
		Message: "Register for a new FtRChat account",
	}

	viewbackend.RegisterMain(w, p)
}

func RegisterMainPost(w http.ResponseWriter, r *http.Request) {
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
	userName := strings.TrimSpace(r.FormValue("auth_name"))

	if userEmail == "" || userPassword == "" || userName == "" {
		http.Error(w, "Email and password are required, but not provided", http.StatusBadRequest)
		return
	}

	fmt.Println("User email:", userEmail)
	fmt.Println("User password:", userPassword)

	db := config.GetDB()

	model_user.CreateUser(db, userName, userEmail, userPassword)

	http.Redirect(w, r, "/success", http.StatusSeeOther)
}
