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

func (s *Server) Start() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
