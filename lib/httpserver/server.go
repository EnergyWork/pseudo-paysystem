package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
	_defaultAddr            = ":9090"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New ...
func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Addr:    _defaultAddr,
		Handler: handler,
		// TLSConfig:         nil,
		ReadTimeout: _defaultReadTimeout,
		// ReadHeaderTimeout: 0,
		WriteTimeout: _defaultWriteTimeout,
		// IdleTimeout:    0,
		// MaxHeaderBytes: 0,
		// TLSNextProto: nil,
		// ConnState:    nil,
		// ErrorLog:    nil,
		// BaseContext: nil,
		// ConnContext: nil,
	}

	srv := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: time.Second * _defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(srv)
	}

	srv.start()

	return srv
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify ...
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown ...
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func (s *Server) Addr() string {
	return s.server.Addr
}
