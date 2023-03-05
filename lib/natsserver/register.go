package natsserver

import (
	"github.com/nats-io/nats.go"
)

var routes map[string]map[string]func() nats.MsgHandler

func init() {
	routes = make(map[string]map[string]func() nats.MsgHandler, 0)
}

func Register(service string, name string, handler func() nats.MsgHandler) {
	if _, ok := routes[service]; !ok {
		routes[service] = make(map[string]func() nats.MsgHandler, 0)
	}
	routes[service][name] = handler
}

func GetHandler(service, name string) func() nats.MsgHandler {
	return routes[service][name]
}

func GetServiceHandlers(service string) map[string]func() nats.MsgHandler {
	return routes[service]
}
