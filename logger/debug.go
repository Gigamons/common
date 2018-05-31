package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// Debug logs a Debug information
func Debug(message string, v ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		if len(v) < 1 {
			log.Println(prefix(color.YellowString("D")), message)
		} else {
			log.Printf(prefix(color.YellowString("D"))+message+"\n", v...)
		}
	}
}
