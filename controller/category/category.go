package category

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SQU1DMAN6/ftrchat/config"
	"github.com/SQU1DMAN6/ftrchat/model"
	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
)

func CategoryNewCategory(w http.ResponseWriter, r *http.Request) {
	paramData := viewbackend.FrontEndParams{
		Title:   "Create a new blog category",
		Message: "Use this page to create a new blog category",
	}

	SS := config.GetSessionManager()

	msg := SS.GetString(r.Context(), "message")
	isLoggedIn := SS.GetBool(r.Context(), "isLoggedIn")
	userName := SS.GetString(r.Context(), "name")

	fmt.Println("msg:", msg)
	fmt.Println("userName:", userName)
	if isLoggedIn != true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	viewbackend.Frontend_CategoryNewCategory(w, paramData)
}

func CategoryNewCategoryPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Now Allowed", http.StatusMethodNotAllowed)
		return
	}

	_ = r.ParseForm

	SS := config.GetSessionManager()
	db := config.GetDB()
	userName := SS.GetString(r.Context(), "name")
	userModel, err := model.GetUserByName(userName, db)
	userID := userModel.ID

	if err != nil {
		http.Error(w, "Failed to identify who you are", http.StatusBadRequest)
		return
	}

	categoryTitle := strings.TrimSpace(r.FormValue("categoryName"))

	if categoryTitle == "" {
		http.Error(w, "The title of the category has not been specified", http.StatusBadRequest)
		return
	}

	fmt.Println("Category title in form:", categoryTitle)

	model.CreateBlogPostCategory(db, categoryTitle, userModel, userID)
}
