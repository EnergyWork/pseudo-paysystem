package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func LoadLogger(dev ...bool) *zerolog.Logger {
	writer := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.StampMicro,
		PartsExclude: []string{
			zerolog.TimestampFieldName,
		},
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		FormatCaller: func(i interface{}) string {
			return filepath.Base(fmt.Sprintf("%s", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("| %s |", i)
		},
		// NoColor:    			false,
		// PartsOrder:          nil,
		// FieldsExclude:       nil, // todo read about that
		// FormatTimestamp:     nil,
		// FormatFieldName:     nil,
		// FormatFieldValue:    nil,
		// FormatErrFieldName:  nil,
		// FormatErrFieldValue: nil,
		// FormatExtra:         nil,
	}
	logger := zerolog.New(writer).With().Timestamp().Caller().Logger()
	if dev[0] == true {
		logger.Level(zerolog.DebugLevel)
	} else {
		logger.Level(zerolog.InfoLevel)
	}
	return &logger
}
