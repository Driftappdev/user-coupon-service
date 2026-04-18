package http

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func New(port string, handler http.Handler) *Server {

	s := &http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{
		server: s,
	}
}

func (s *Server) Start() {

	go func() {
		log.Println("HTTP started on", s.server.Addr)

		if err := s.server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			log.Fatal(err)
		}
	}()
}

func (s *Server) Stop(ctx context.Context) error {

	return s.server.Shutdown(ctx)
}

