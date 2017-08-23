package logger

import (
	"fmt"
)

var gDefaultLog *Logger

func init() {
	var err error
	gDefaultLog, err = NewLogger(NewDefaultConfig())
	if err != nil {
		fmt.Println(err)
	}
}

func SetDefaultLogger(logger *Logger) {
	gDefaultLog = logger
}

func Debug(v ...interface{}) {
	if gDefaultLog == nil {
		fmt.Println(NewFormat(DEBUG, v, 2).ColorString())
		return
	}

	gDefaultLog.Log(1, DEBUG, v)
}

func Info(v ...interface{}) {
	if gDefaultLog == nil {
		fmt.Println(NewFormat(INFO, v, 2))
		return
	}

	gDefaultLog.Log(1, INFO, v)
}

func Warn(v ...interface{}) {
	if gDefaultLog == nil {
		fmt.Println(NewFormat(WARN, v, 2))
		return
	}

	gDefaultLog.Log(1, WARN, v)
}

func Error(v ...interface{}) {
	if gDefaultLog == nil {
		fmt.Println(NewFormat(ERROR, v, 2))
		return
	}

	gDefaultLog.Log(1, ERROR, v)
}

func Fatal(v ...interface{}) {
	if gDefaultLog == nil {
		fmt.Println(NewFormat(FATAL, v, 2))
		return
	}

	gDefaultLog.Log(1, FATAL, v)
}

func Log(skip int, level Level, v []interface{}) {
	if gDefaultLog == nil {
		fmt.Println(NewFormat(level, v, skip+2))
		return
	}

	gDefaultLog.Log(skip, level, v)
}

func GetDefaultLogger() *Logger {
	return gDefaultLog
}
