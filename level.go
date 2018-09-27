package logger

import (
	"fmt"
)

type Level int

func (l Level) String() string {
	return levelStringMap[l]
}

func (l Level) ConsoleColorString() string {
	return fmt.Sprintf("\033[;%dm%s\033[0m", l.ConsoleColorNum(), l.String())
}

func (l Level) ConsoleColorNum() int {
	switch l {
	case DEBUG:
		return 34
	case INFO:
		return 32
	case WARN:
		return 33
	case ERROR:
		return 35
	case FATAL:
		return 31
	default:
		return 37
	}
}

var levelStringMap = []string{
	"ALL", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "OFF",
}

const (
	ALL Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)
