package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

// Info logs an informational message
func Info(args ...interface{}) {
	log.Info(args...)
}

// Infof logs an informational message with formatting
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Error logs an error message
func Error(args ...interface{}) {
	log.Error(args...)
}

// Errorf logs an error message with formatting
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatal logs a fatal error message and exits the application
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf logs a fatal error message with formatting and exits the application
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
