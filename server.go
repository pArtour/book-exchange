package go_booc_exchange

import (
	"context"
	"github.com/pArtour/book-exchange/configs"
	"net/http"
)

type Server struct {
	config     *configs.Config
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + s.config.Server.BindAddr,
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx *context.Context) error {
	return s.Shutdown(ctx)
}

func NewServer(config *configs.Config) *Server {
	return &Server{config: config}
}
