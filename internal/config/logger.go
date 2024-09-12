package config

import (
	"os"
)

const (
	logLevelEnv  = "TESTER_LOG_LEVEL"
	logFormatEnv = "TESTER_LOG_FORMAT"
)

const (
	defaultLevel  = "info"
	defaultFormat = "15:04:05"
)

// LogConfig defines the interface for configuring logging settings.
type LogConfig interface {
	Level() string
	Format() string
}

// NewLoggerConfig creates and returns a new instance of LogConfig.
func NewLoggerConfig() (LogConfig, error) {
	level := os.Getenv(logLevelEnv)
	if level == "" {
		level = defaultLevel
	}

	format := os.Getenv(logFormatEnv)
	if format == "" {
		format = defaultFormat
	}

	l := loggerConfig{
		level:  level,
		format: format,
	}

	return &l, nil
}

type loggerConfig struct {
	level  string
	format string
}

func (l *loggerConfig) Level() string {
	return l.level
}

func (l *loggerConfig) Format() string {
	return l.format
}
