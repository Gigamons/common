package logger

import (
	"os"
	"testing"
)

func TestDebugln(t *testing.T) {
	t.Parallel()
	os.Setenv("DEBUG", "false")
	Debugln("Sample Debug message, that message shouldn't show.")
	Debugln() // Should not print [Debug]
	os.Setenv("DEBUG", "true")
	Debugln("Sample Debug message, that message should show!")
	Debugln() // Should only print [Debug]
}

func TestDebugf(t *testing.T) {
	t.Parallel()
	os.Setenv("DEBUG", "false")
	Debugln("Sample Debug message, that message shouldn't show.")
	Debugln() // Should not print [Debug]
	os.Setenv("DEBUG", "true")
	Debugln("Sample Debug message, that message should show!")
	Debugln() // Should only print [Debug]
}
