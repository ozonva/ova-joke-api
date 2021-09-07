package logger

import (
	"github.com/rs/zerolog/log"
)

func Infof(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func Warnf(format string, v ...interface{}) {
	log.Warn().Msgf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	log.Error().Msgf(format, v...)
}
