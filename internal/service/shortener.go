package service

import "net/http"

type Shortener struct {
	server *http.Server
}

func NewShortener(h http.Handler, addr string) *Shortener {
	return &Shortener{
		server: &http.Server{
			Handler: h,
			Addr:    addr,
		},
	}
}

func (s *Shortener) Start() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
