package server

import "net/http"

type Server struct {
	server *http.Server
}

func New(h http.Handler, addr string) *Server {
	httpServer := &http.Server{
		Handler: h,
	}

	s := &Server{
		server: httpServer,
	}

	s.server.Addr = addr

	return s
}

func (s *Server) Start() {
	s.server.ListenAndServe()
}
