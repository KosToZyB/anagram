package server

import (
	"net/http"

	"github.com/kostozyb/anagram/internal/config"
	"github.com/kostozyb/anagram/internal/handlers"
	"github.com/kostozyb/anagram/pkg/logger"
)

type Server struct {
	log logger.Logger
	cfg config.Configer
}

func NewServer(cfg config.Configer, log logger.Logger) *Server {
	return &Server{
		log: log,
		cfg: cfg,
	}
}

func (s *Server) Run() {
	api := handlers.NewAPI(s.log)

	hs := &http.Server{
		Addr:    ":8080",
		Handler: api.Router(),
	}

	if err := hs.ListenAndServe(); err != nil {
		s.log.Warnf("server stopped or not started: %s", err)
	}
}