package logger

import (
	"log"

	"github.com/fatih/color"
)

// Infof formats and sends an Info to std out or file to log.
func Infof(format string, v ...interface{}) {
	log.Printf(prefix(color.CyanString("I"))+" "+format, v...)
}

// Infoln sends an Info to std out or file to log.
func Infoln(v ...interface{}) {
	v = append(ToInterface(prefix(color.CyanString("I"))), v...)
	log.Println(v...)
}
