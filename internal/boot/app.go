package boot

import (
	"net/http"

	routes "github.com/SQU1DMAN6/ftrchat/internal/module"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func BootApp() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	routes.RegisterRoutes(r)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello route welcome"))
	})
	http.ListenAndServe(":3000", r)
}
