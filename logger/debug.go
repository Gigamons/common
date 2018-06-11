package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// Debugf logs an Debug message. (IF Env DEBUG == true)
func Debugf(format string, v ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		log.Printf(prefix(color.BlueString("Debug"))+" "+format, v...)
	}
}

// Debugln prints an Debug with a new line at end. (IF Env DEBUG == true)
func Debugln(v ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		v = append(ToInterface(prefix(color.BlueString("Debug"))), v...)
		log.Println(v...)
	}
}
