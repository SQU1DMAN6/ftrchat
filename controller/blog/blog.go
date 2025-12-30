package blog

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SQU1DMAN6/ftrchat/config"
	"github.com/SQU1DMAN6/ftrchat/model"
	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
)

func BlogNewBlog(w http.ResponseWriter, r *http.Request) {
	paramData := viewbackend.FrontEndParams{
		Title:   "Create a new blog",
		Message: "Use this page to write a new blog post.",
	}

	SS := config.GetSessionManager()

	msg := SS.GetString(r.Context(), "message")
	isLoggedIn := SS.GetBool(r.Context(), "isLoggedIn")
	userName := SS.GetString(r.Context(), "name")

	fmt.Println("msg:", msg)
	fmt.Println("userName", userName)
	if isLoggedIn != true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	viewbackend.Frontend_BlogNewBlog(w, paramData)
}

func BlogListBlogs(w http.ResponseWriter, r *http.Request) {
	paramData := viewbackend.FrontEndParams{
		Title:   "FtR blog",
		Message: "Use this page to view existing posts.",
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

	viewbackend.Frontend_BlogMain(w, paramData)
}

func BlogNewBlogPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	_ = r.ParseForm

	SS := config.GetSessionManager()
	db := config.GetDB()
	userName := SS.GetString(r.Context(), "name")
	userModel, errasasad := model.GetUserByName(userName, db)
	userID := userModel.ID
	if errasasad != nil {
		http.Error(w, "Failed to identify who you are", http.StatusBadRequest)
		return
	}

	blogTitle := strings.TrimSpace(r.FormValue("title"))
	blogContents := strings.TrimSpace(r.FormValue("blogContents"))

	if blogTitle == "" || blogContents == "" {
		http.Error(w, "Blog title and contents are required", http.StatusBadRequest)
		return
	}

	fmt.Println("Blog title in form:", blogTitle)
	fmt.Println("Blog contents in form:", blogContents)

	timestamp := time.Now().Unix()

	model.NewBlogPost(db, blogTitle, "GENERAL", timestamp, userID)
}
