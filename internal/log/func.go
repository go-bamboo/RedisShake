package log

import (
	"github.com/go-bamboo/pkg/log"
)

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
