package main

import (
    "log"
    "net/http"
)

// HookFunc defines the signature for error callbacks.
type HookFunc func(error)

// Option applies configuration to Server.
type Option func(*Server)

// Server is a minimal HTTP server with an OnError hook.
type Server struct {
    mux     *http.ServeMux
    onError HookFunc
}

// WithOnError sets a custom error hook.
func WithOnError(h HookFunc) Option {
    return func(s *Server) { s.onError = h }
}

// NewServer constructs a Server with default no‑op hook.
func NewServer(opts ...Option) *Server {
    s := &Server{
        mux:     http.NewServeMux(),
        onError: func(error) {}, // no‑op
    }
    for _, opt := range opts {
        opt(s)
    }
    s.routes()
    return s
}

func (s *Server) routes() {
    s.mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello"))
    })
}

// Listen starts serving.
func (s *Server) Listen(addr string) error {
    log.Printf("Listening on %s", addr)
    return http.ListenAndServe(addr, s.mux)
}
