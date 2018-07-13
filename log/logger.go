package golog

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/Silverhammertech/sms-svc/config"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

func Debug(msg ...interface{}) {
	if config.DEBUG {
		log.Debug(msg...)
	}
}

func Info(msg ...interface{}) {
	log.Info(msg...)
}

func Warn(msg ...interface{}) {
	log.Warn(msg...)
}

func Error(msg ...interface{}) {
	log.Error(msg...)
}

func Fatal(msg ...interface{}) {
	// Calls os.Exit(1) after logging
	log.Fatal(msg...)
}

func Panic(msg ...interface{}) {
	log.Panic(msg...)
}