package routes

import (
	"net/http"

	"github.com/SQU1DMAN6/ftrchat/controller/index"
	"github.com/SQU1DMAN6/ftrchat/controller/login"
	"github.com/SQU1DMAN6/ftrchat/controller/something"

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
}
