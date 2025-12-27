package routes

import (
	"net/http"

	"github.com/SQU1DMAN6/ftrchat/controller/blog"
	"github.com/SQU1DMAN6/ftrchat/controller/chat"
	"github.com/SQU1DMAN6/ftrchat/controller/index"
	"github.com/SQU1DMAN6/ftrchat/controller/login"
	"github.com/SQU1DMAN6/ftrchat/controller/register"
	"github.com/SQU1DMAN6/ftrchat/controller/something"
	"github.com/SQU1DMAN6/ftrchat/controller/success"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello route welcome"))
	})
	r.Get("/quan", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello route localhost:6769/quan welcome"))
	})
	r.Get("/", index.Index)
	r.Get("/other", index.Other)
	r.Get("/ftr", something.Whatever)
	r.Get("/login", login.LoginMain)
	r.Post("/login", login.LoginMainPost)
	r.Get("/register", register.RegisterMain)
	r.Post("/register", register.RegisterMainPost)
	r.Get("/success", success.SuccessRegister)
	r.Get("/newblog", blog.BlogNewBlog)
	r.Get("/blog", blog.BlogListBlogs)
	r.Post("/newblog", blog.BlogNewBlogPost)

	hub := chat.NewHub()
	go hub.Run()
	r.Get("/chat", chat.ChatMain)
	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})
}
