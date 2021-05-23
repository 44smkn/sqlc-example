package presentation

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

var log *zap.Logger

func ApiRouter() http.Handler {
	r := chi.NewRouter()
	r.Route("/chair", func(r chi.Router) {
		r.Get("/{id}", getChairDetail)
		r.Post("/", postChair)
		r.Get("/search", searchChair)
	})

	return r
}

func InitLogger(logger *zap.Logger) {
	log = logger
}
