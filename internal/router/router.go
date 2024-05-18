package router

import (
	"net/http"
)

type ServerMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Router struct {
	Mux ServerMux

	pattern string
	handler func(http.ResponseWriter, *http.Request)
}

func NewMuxRouter(m ServerMux) *Router {
	return &Router{
		Mux: m,
	}
}

func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) *Router {
	r.pattern = pattern
	r.handler = handler

	r.Mux.HandleFunc(r.pattern, r.handler)

	return r
}

func (r *Router) WithTrailingSlash() {
	r.Mux.HandleFunc(r.pattern+"/{$}", r.handler)
}
