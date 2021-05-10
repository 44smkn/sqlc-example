package presentation

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Mount("/api", apiRouter())

	return r
}

func apiRouter() http.Handler {
	r := chi.NewRouter()
	r.Route("/chair", func(r chi.Router) {
		r.Get("/{id}", getChairDetail)
		r.Post("/", postChair)
		r.Get("/search", searchChair)
	})

	return r
}
