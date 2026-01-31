package blog

import (
	"fmt"
	"html/template"
	"math"
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

	offset := 5

	pageIDInt, err := strconv.ParseInt(pageID, 10, 64)
	if pageID != "next" && pageID != "previous" && err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse page ID: %s", err), http.StatusBadRequest)
		return
	}

	blogPostsData, err := model.ListBlogPostsWithPagination(db, int(pageIDInt), offset)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get blog post listing: %s", err), http.StatusBadRequest)
		return
	}

	fmt.Println("==================")

	fmt.Println("bloglog blog", blogPostsData)

	fmt.Println("==================")

	blogPosts := blogPostsData.Posts
	previousblogPosts := blogPostsData.PreviousPage
	nextblogPosts := blogPostsData.NextPage
	totalNumberOfPosts := blogPostsData.TotalPost

	paramData.Pagination = make(map[string]interface{})
	paramData.Pagination["Previous"] = previousblogPosts
	paramData.Pagination["Next"] = nextblogPosts
	paramData.Pagination["TotalPages"] = int(math.Ceil(float64(float64(totalNumberOfPosts) / float64(offset))))
	fmt.Println("================================")
	fmt.Println(blogPostsData)
	fmt.Println("================================")
	fmt.Println("Total number of posts:", totalNumberOfPosts)
	fmt.Println("Total number of pages:", paramData.Pagination)

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
	// blogCategoryRaw := r.FormValue("category")
	// blogCategory, err := strconv.ParseInt(blogCategoryRaw, 10, 0)

	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Failed to pass blog category as integer: %s", err), http.StatusBadRequest)
	// 	return
	// }

	if blogTitle == "" || blogContents == "" {
		http.Error(w, "Blog title and contents are required", http.StatusBadRequest)
		return
	}

	fmt.Println("Blog title in form:", blogTitle)
	fmt.Println("Blog contents in form:", blogContents)

	timestamp := time.Now().Unix()

	// var category *model.BlogCategory

	// category = &model.BlogCategory{Name: "GENERAL", User: userModel, UserID: userModel.ID}

	id, err := model.NewBlogPost(db, blogTitle, blogContents, 1, timestamp, userID)
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

	var postID int64
	var err error
	postIDString := chi.URLParam(r, "pid")
	postID, err = strconv.ParseInt(postIDString, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse post ID: %s", err), http.StatusBadRequest)
		return
	}

	fmt.Printf("User %s tried to access post %d\n", userName, postID)

	blogPostModel, err := model.GetBlogPost(int(postID), db)

	// fmt.Println("blogpost, hopefully with user info", blogPostModel)

	// fmt.Println("blogpost, hopefully with user info, that is Name: ", blogPostModel.User.Name)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve blog post: %s", err), http.StatusBadRequest)
		return
	}

	blogContentsCooked := strings.ReplaceAll(blogPostModel.Contents, "\n", "<br style='line-height: 1'>")

	paramData := viewbackend.FrontEndParams{
		Title:    blogPostModel.Title,
		Name:     blogPostModel.User.Name,
		SafeBody: template.HTML(blogContentsCooked),
	}

	viewbackend.Frontend_BlogView(w, paramData)

	// Use this id to query to article/post
	//SELECT * FROM articles/blog WHERE ID = postId

	// ctx := r.Context()
	// key := ctx.Value("key").(string)
}
