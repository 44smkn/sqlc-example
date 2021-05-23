package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/44smkn/sqlc-sample/pkg/presentation"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type Server struct {
	Port   int16
	Logger *zap.Logger
}

func NewServer(port int16, logger *zap.Logger) *Server {
	return &Server{
		Port:   port,
		Logger: logger,
	}
}

func (s *Server) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	presentation.InitLogger(s.Logger)
	router := newRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", s.Port),
		Handler: router,
	}
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	return server.ListenAndServe()
}

func newRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})
	r.Mount("/api", presentation.ApiRouter())

	return r
}
