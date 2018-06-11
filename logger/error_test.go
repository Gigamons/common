package logger

import (
	"errors"
	"testing"
)

func TestErrorln(t *testing.T) {
	t.Parallel()
	Errorln("Sample error message.", errors.New("Sample Error"))
	Errorln()
	Errorln("Test")
}

func TestErrorf(t *testing.T) {
	t.Parallel()
	Errorf("Sample Test %v%s%v", 10, "Test", errors.New("Testing"))
	Errorf("Test123")
}
