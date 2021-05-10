package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/44smkn/sqlc-sample/pkg/presentation"
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

	router := presentation.NewRouter()
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
