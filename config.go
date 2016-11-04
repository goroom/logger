package logger

import (
	"strings"
)

type Unit int64

const (
	_       = iota
	KB Unit = 1 << (iota * 10)
	MB
	GB
	TB
)

type Level int

const (
	ALL Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

func (level Level) String() string {
	switch level {
	case ALL:
		return "ALL"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	case OFF:
		return "OFF"
	}
	return ""
}

func StringLevel(level string) Level {
	level = strings.ToUpper(level)
	switch level {
	case "ALL":
		return ALL
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return OFF
	}
}

type Config struct {
	ConsoleLevel    Level
	FilePath        string
	FileBaseName    string
	Level           Level
	MaxFileCount    int
	MaxFileSize     int64
	MaxFileSizeUnit Unit
	CallBackFunc    func(format *Format)
}

func NewDefaultConfig() *Config {
	var config Config
	config.ConsoleLevel = ALL
	config.FilePath = "./"
	config.FileBaseName = "main.log"
	config.Level = ALL
	config.MaxFileCount = 5
	config.MaxFileSize = 5
	config.MaxFileSizeUnit = MB

	return &config
}
