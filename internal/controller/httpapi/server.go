package httpapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gmcriptobox/otus-go-final-project/internal/config"
)

type Server struct {
	server  *http.Server
	handler http.Handler
	config  *config.Config
}

func NewServer(handler http.Handler, config *config.Config) *Server {
	return &Server{
		config:  config,
		handler: handler,
	}
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:         s.config.Server.Port,
		Handler:      s.handler,
		ReadTimeout:  time.Duration(s.config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.Server.WriteTimeout) * time.Second,
	}
	return s.server.ListenAndServe()
}

func (s *Server) ShutdownService(ctx context.Context, cancel context.CancelFunc) {
	<-ctx.Done()
	s.server.Shutdown(ctx)
	cancel()
}
