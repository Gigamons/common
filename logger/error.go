package logger

import (
	"log"

	"github.com/fatih/color"
)

// Errorf logs an Error if even.
func Errorf(format string, v ...interface{}) {
	log.Printf(prefix(color.RedString("ERR"))+" "+format, v...)
}

// Errorln prints an Error + add a new Line.
func Errorln(v ...interface{}) {
	v = append(ToInterface(prefix(color.RedString("ERR"))), v...)
	log.Println(v...)
}
