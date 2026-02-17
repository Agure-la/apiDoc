package http

import (
	"context"
	"log"
	"net/http"

	"github.com/agure-la/api-docs/internal/config"
	"github.com/agure-la/api-docs/internal/spec"
)

type Server struct {
	config    *config.Config
	service   *spec.Service
	router    *Router
	httpServer *http.Server
}

func NewServer(cfg *config.Config) *Server {
	service := spec.NewService(cfg)
	router := NewRouter(service)

	return &Server{
		config:  cfg,
		service: service,
		router:  router,
	}
}

func (s *Server) Start() error {
	// Load all specifications on startup
	if err := s.service.LoadAll(); err != nil {
		return err
	}

	s.httpServer = &http.Server{
		Addr:         ":" + s.config.Server.Port,
		Handler:      s.router.SetupRoutes(),
		ReadTimeout:  s.config.Server.ReadTimeout,
		WriteTimeout: s.config.Server.WriteTimeout,
		IdleTimeout:  s.config.Server.IdleTimeout,
	}

	log.Printf("Server starting on port %s", s.config.Server.Port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server...")
	return s.httpServer.Shutdown(ctx)
}
