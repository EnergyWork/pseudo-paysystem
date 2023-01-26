package logger

import (
	"fmt"

	"github.com/rs/zerolog"
)

type CLogger struct {
	prefix string
	logger *zerolog.Logger
}

func New(l *zerolog.Logger) *CLogger {
	return &CLogger{logger: l}
}

func (l *CLogger) Info(m ...interface{}) {
	l.logger.Info().Msg(l.prefix + parseMsg(m...))
}

func (l *CLogger) Warn(m ...interface{}) {
	l.logger.Warn().Msg(l.prefix + parseMsg(m...))
}

func (l *CLogger) Error(m ...interface{}) {
	l.logger.Error().Msg(l.prefix + parseMsg(m...))
}

func (l *CLogger) Fatal(m ...interface{}) {
	l.logger.Fatal().Msg(l.prefix + parseMsg(m...))
}

func (l *CLogger) Panic(m ...interface{}) {
	l.logger.Panic().Msg(l.prefix + parseMsg(m...))
}

func (l *CLogger) Debug(m ...interface{}) {
	l.logger.Debug().Msg(l.prefix + parseMsg(m...))
}

func (l *CLogger) Trace(m ...interface{}) {
	l.logger.Trace().Msg(l.prefix + parseMsg(m...))
}

func (l *CLogger) WithPrefix(prefix string) *CLogger {
	var p string
	if len(prefix) > 0 {
		p = prefix // + ": "
	}
	if len(p) > 0 {
		l.prefix = p + ": "
	}
	return l
}

func parseMsg(m ...interface{}) string {
	if len(m) == 0 {
		return ""
	} else if len(m) == 1 {
		t := m[0]
		s, ok := t.(string)
		if ok {
			return s
		} else {
			return ""
		}
	} else {
		f, ok := m[0].(string)
		if ok {
			return fmt.Sprintf(f, m[1:]...)
		} else {
			return ""
		}
	}
}
