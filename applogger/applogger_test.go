package applogger

import (
	"os"
	"testing"
)

func deleteLogFile(t *testing.T) func() {
	return func() {
		f := "applogger.log"
		_, err := os.Stat(f)
		if err == nil {
			_ = os.Remove(f)
		}
	}
}

func TestGetInstance(t *testing.T) {
	deleteLogFile(t)
	logger := GetInstance()
	if logger == nil {
		t.Error("logger is null.")
	}

	logger2 := GetInstance()
	if logger.createAt != logger2.createAt {
		t.Error("loggers are different instanance.")
	}
}
