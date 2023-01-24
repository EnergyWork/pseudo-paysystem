package domain

type Logger interface {
	Info(...interface{})
	Debug(...interface{})
	Warn(...interface{})
	Fatal(...interface{})
	Error(...interface{})
	Panic(...interface{})
	Trace(...interface{})
}
