package natsserver

import (
	"github.com/nats-io/nats.go"
)

type Server struct {
	nc      *nats.Conn
	subject string
	notify  chan error
}

func New(nc *nats.Conn, queue string, handlers map[string]func() nats.MsgHandler) (*Server, error) {
	srv := &Server{nc: nc}

	for subject, handler := range handlers {
		sub, err := srv.nc.QueueSubscribe(subject, queue, handler())
		if err != nil || sub == nil {
			srv.notify <- err

		}
	}
	if err := srv.nc.Flush(); err != nil {
		srv.notify <- err
	}

	return srv, nil // error?
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	return s.nc.Drain()
}
