package routes

import (
	"github.com/SQU1DMAN6/ftrchat/internal/module/index"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Get("/", index.Index)
	r.Get("/other", index.Other)
}
