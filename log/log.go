package log

import (
	"fmt"
	"os"

	zl "github.com/rs/zerolog"
	l "github.com/rs/zerolog/log"
)

func init() {
	levelString := os.Getenv("LOG_LEVEL")
	level, err := zl.ParseLevel(levelString)
	if err != nil {
		fmt.Printf("Unable to parse log level [%s] with error [%s]. Defaulting to level INFO", levelString, err)
		level = zl.InfoLevel
	}
	zl.SetGlobalLevel(level)
	l.Logger = l.Output(zl.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: " ",
	})
}
func Debugf(format string, args ...interface{}) { l.Debug().Msgf(format, args...) }
func Infof(format string, args ...interface{})  { l.Info().Msgf(format, args...) }
func Warnf(format string, args ...interface{})  { l.Warn().Msgf(format, args...) }
func Errorf(format string, args ...interface{}) { l.Error().Msgf(format, args...) }
func Fatalf(format string, args ...interface{}) { l.Fatal().Msgf(format, args...) }
