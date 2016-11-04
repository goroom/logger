package logger

import (
	"fmt"
)

var g_default_log *Logger

func SetDefaultLogger(logger *Logger) {
	g_default_log = logger
}

func Debug(v ...interface{}) {
	if g_default_log == nil {
		fmt.Println(NewFormat(DEBUG, v, 2))
		return
	}

	g_default_log.Log(1, DEBUG, v)
}

func Info(v ...interface{}) {
	if g_default_log == nil {
		fmt.Println(NewFormat(INFO, v, 2))
		return
	}

	g_default_log.Log(1, INFO, v)
}

func Warn(v ...interface{}) {
	if g_default_log == nil {
		fmt.Println(NewFormat(WARN, v, 2))
		return
	}

	g_default_log.Log(1, WARN, v)
}

func Error(v ...interface{}) {
	if g_default_log == nil {
		fmt.Println(NewFormat(ERROR, v, 2))
		return
	}

	g_default_log.Log(1, ERROR, v)
}

func Fatal(v ...interface{}) {
	if g_default_log == nil {
		fmt.Println(NewFormat(FATAL, v, 2))
		return
	}

	g_default_log.Log(1, FATAL, v)
}

func Log(skip int, level Level, v []interface{}) {
	if g_default_log == nil {
		fmt.Println(NewFormat(level, v, skip+2))
		return
	}

	g_default_log.Log(skip, level, v)
}

func GetDefaultLogger() *Logger {
	return g_default_log
}
