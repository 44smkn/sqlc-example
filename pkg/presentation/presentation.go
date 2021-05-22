package presentation

import (
	"net/http"

	"github.com/go-chi/chi"
)

func ApiRouter() http.Handler {
	r := chi.NewRouter()
	r.Route("/chair", func(r chi.Router) {
		r.Get("/{id}", getChairDetail)
		r.Post("/", postChair)
		r.Get("/search", searchChair)
	})

	return r
}
