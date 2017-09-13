package logger

import (
	"fmt"
	"testing"
)

type A struct {
	a int
	b string
}

func TestA(t *testing.T) {
	default_config := NewDefaultConfig()
	default_config.FilePath = "log"
	default_config.FileBaseName = "logger.log"
	err := Init(default_config)
	if err != nil {
		t.Error(err)
		return
	}
	SetMaxFileSize(1, KB)
	Debug("debug_message", A{a: 1, b: "2"})

	SetFileLevel(INFO)
	SetConsoleLevel(INFO)
	Debug("debug_message")

	Info("info_message")
	Warn("warn_message")

	SetCallBackFunc(func(f *Format) {
		fmt.Println("Call back", f)
	})

	Error("error_message")
	Fatal("fatal_message")
	Debug("debug_message", 2)

	SetConsoleLevel(OFF)
	SetFileLevel(OFF)

	Error("detault_debug_message")

	//time.Sleep(1e9)
}
