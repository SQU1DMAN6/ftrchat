package app

import (
	"net/http"

	routes "github.com/SQU1DMAN6/ftrchat"
	"github.com/go-chi/chi/v5"
)

func BootApp() {
	r := chi.NewRouter()
	RegisterMiddlewares(r)
	routes.RegisterRoutes(r)

	http.ListenAndServe(":6769", r)
}
