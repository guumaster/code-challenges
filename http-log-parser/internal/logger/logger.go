package logger

import (
	"github.com/sirupsen/logrus"
)

// NewLogger returns an instance of logrus.
func NewLogger(level LogLevel) Logger {
	log := logrus.New()

	var logrusLevel logrus.Level

	switch level {
	case Debug:
		logrusLevel = logrus.DebugLevel
	case Info:
		logrusLevel = logrus.InfoLevel
	case Warning:
		logrusLevel = logrus.WarnLevel
	case Error:
		logrusLevel = logrus.ErrorLevel
	default:
		logrusLevel = logrus.InfoLevel
	}

	log.SetLevel(logrusLevel)

	return log
}

// LogLevel string type to indicate common log levels.
type LogLevel string

// Debug debug level.
const Debug LogLevel = "debug"

// Info info level.
const Info LogLevel = "info"

// Warning warning level.
const Warning LogLevel = "warning"

// Error error level.
const Error LogLevel = "error"

// Logger common interface.
type Logger interface {
	Fatalf(string, ...interface{})
	Errorf(string, ...interface{})
	Warningf(string, ...interface{})
	Infof(string, ...interface{})
	Debugf(string, ...interface{})
}
