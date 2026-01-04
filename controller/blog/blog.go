package blog

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/SQU1DMAN6/ftrchat/config"
	"github.com/SQU1DMAN6/ftrchat/model"
	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
	"github.com/go-chi/chi/v5"
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
	db := config.GetDB()

	pageID := chi.URLParam(r, "pid")
	if pageID == "" {
		http.Redirect(w, r, "/blog/1", http.StatusSeeOther)
		return
	}
	pageIDInt, err := strconv.ParseInt(pageID, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse page ID: %s", err), http.StatusBadRequest)
		return
	}

	blogPosts, err := model.ListBlogPostsWithPagination(db, int(pageIDInt))

	fmt.Println("================================")
	fmt.Println(blogPosts)
	fmt.Println("================================")

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get blog post listing: %s", err), http.StatusInternalServerError)
		return
	}

	paramData.UserData = blogPosts

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

	id, err := model.NewBlogPost(db, blogTitle, blogContents, "GENERAL", timestamp, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create new post: %s", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/view/%d", id), http.StatusSeeOther)
}

func BlogViewBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	SS := config.GetSessionManager()
	isLoggedIn := SS.GetBool(r.Context(), "isLoggedIn")

	if isLoggedIn != true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	db := config.GetDB()

	userName := SS.GetString(r.Context(), "name")

	postID := chi.URLParam(r, "pid")
	postIDInt, err := strconv.ParseInt(postID, 10, 64)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse post ID: %s", err), http.StatusBadRequest)
		return
	}

	fmt.Printf("User %s tried to access post %d\n", userName, postIDInt)

	blogPostModel, err := model.GetBlogPost(int(postIDInt), db)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve blog post: %s", err), http.StatusBadRequest)
		return
	}

	blogContentsCooked := strings.ReplaceAll(blogPostModel.Contents, "\n", "<br style='line-height: 1'>")

	paramData := viewbackend.FrontEndParams{
		Title:    blogPostModel.Title,
		SafeBody: template.HTML(blogContentsCooked),
	}

	viewbackend.Frontend_BlogView(w, paramData)

	// Use this id to query to article/post
	//SELECT * FROM articles/blog WHERE ID = postId

	// ctx := r.Context()
	// key := ctx.Value("key").(string)
}
