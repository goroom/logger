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
	log, err := NewLogger(default_config)
	if err != nil {
		t.Error(err)
		return
	}
	log.SetMaxFileSize(1, KB)
	log.Debug("debug_message", A{a: 1, b: "2"})

	log.SetLevel(INFO)
	log.SetConsoleLevel(INFO)
	log.Debug("debug_message")

	log.Info("info_message")
	log.Warn("warn_message")

	log.SetCallBackFunc(func(f *Format) {
		fmt.Println("Call back", f)
	})

	log.Error("error_message")
	log.Fatal("fatal_message")

	SetDefaultLogger(log)
	Debug("debug_message", 2)

	log.SetConsoleLevel(OFF)
	log.SetLevel(OFF)

	Error("detault_debug_message")
}
