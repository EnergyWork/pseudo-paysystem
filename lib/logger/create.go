package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoadLogger(dev ...bool) *zerolog.Logger {
	writer := zerolog.ConsoleWriter{
		Out: os.Stderr,
		// TimeFormat: time.StampMicro,
		/*PartsExclude: []string{
			zerolog.TimestampFieldName,
		},*/
		FormatTimestamp: func(i interface{}) string {
			return time.Now().Format(time.StampMicro)
		},
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		/*FormatCaller: func(i interface{}) string {
			return filepath.Base(fmt.Sprintf("%s", i))
		},*/
		FormatMessage: func(i interface{}) string {
			if msg, ok := i.(string); ok && len(msg) > 1000 {
				return fmt.Sprintf("|SHORT: %s... |", i.(string)[:200])
			} else {
				return fmt.Sprintf("| %s |", i)
			}
		},
		// Additional parameters
		// NoColor: false,
		// PartsOrder:          nil,
		// FieldsExclude:       nil, // todo read about that
		// FormatTimestamp:     nil,
		// FormatFieldName:     nil,
		// FormatFieldValue:    nil,
		// FormatErrFieldName:  nil,
		// FormatErrFieldValue: nil,
		// FormatExtra:         nil,
		// PartsExclude			nil,
	}
	logger := log.Output(writer)

	if dev[0] == true {
		logger.Level(zerolog.DebugLevel)
	} else {
		logger.Level(zerolog.InfoLevel)
	}

	return &logger
}
