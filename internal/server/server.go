package server

import "net/http"

type Server struct {
	server *http.Server
}

func New(h http.Handler) *Server {
	httpServer := &http.Server{
		Handler: h,
	}

	s := &Server{
		server: httpServer,
	}

	s.server.Addr = "localhost:8080"

	return s
}

func (s *Server) Start() {
	s.server.ListenAndServe()
}
