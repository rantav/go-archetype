package log

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	logger zerolog.Logger
}

func NewZeroLogger(level string) *ZeroLogger {
	zlLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to parse log level [%s] with error [%s]. Defaulting to level INFO", level, err)
		zlLevel = zerolog.InfoLevel
	}

	return &ZeroLogger{
		logger: zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: " ",
		}).With().Timestamp().Logger().Level(zlLevel),
	}
}

func (l *ZeroLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

func (l *ZeroLogger) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

func (l *ZeroLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

func (l *ZeroLogger) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}

func (l *ZeroLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
}
