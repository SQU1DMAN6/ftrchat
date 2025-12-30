package blog

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SQU1DMAN6/ftrchat/config"
	"github.com/SQU1DMAN6/ftrchat/model"
	viewbackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
	"github.com/go-chi/chi/v5"
)

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

	// html to view the blog

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

	paramData := viewbackend.FrontEndParams{
		Title:    blogPostModel.Title,
		Message:  blogPostModel.Contents,
		SafeBody: "<b> hello strong html tg </b> <h1> this is h1 </h1>",
	}

	viewbackend.Frontend_BlogView(w, paramData)

	// Use this id to query to article/post
	//SELECT * FROM articles/blog WHERE ID = postId

	// ctx := r.Context()
	// key := ctx.Value("key").(string)
}
