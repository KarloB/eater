package server

import "net/http"

// Server server
type Server struct {
	Handlers map[string]Handler
}

// Handler handler
type Handler struct {
	Path string
	Func http.Handler
}

// New init server
func New() *Server {
	t := Server{}

	handlers := map[string]Handler{
		"eater": Handler{Path: "/", Func: http.HandlerFunc(t.Eater)},
	}
	t.Handlers = handlers

	return &t
}
