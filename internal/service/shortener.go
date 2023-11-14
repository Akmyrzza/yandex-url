package service

import "net/http"

type Shortener struct {
	Server *http.Server
}

func NewShortener(h http.Handler, addr string) *Shortener {
	return &Shortener{
		Server: &http.Server{
			Handler: h,
			Addr:    addr,
		},
	}
}
