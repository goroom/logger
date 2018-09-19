package logger

import (
	"context"
	"fmt"
)

var defaultLogger *Logger

func init() {
	defaultLogger = newLogger("")
	fmt.Println("default log file", defaultLogger.config.GetFilePath())
}

func GetDefaultLogger() *Logger {
	return defaultLogger
}

func Debug(args ...interface{}) {
	defaultLogger.log(1, DEBUG, args)
}

func DebugF(format string, args ...interface{}) {
	defaultLogger.logF(1, DEBUG, format, args)
}

func CDebug(ctx context.Context, args ...interface{}) {
	defaultLogger.cLog(ctx, 1, DEBUG, args)
}

func CDebugF(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.cLogF(ctx, 1, DEBUG, format, args)
}

func Info(args ...interface{}) {
	defaultLogger.log(1, INFO, args)
}

func InfoF(format string, args ...interface{}) {
	defaultLogger.logF(1, INFO, format, args)
}

func CInfo(ctx context.Context, args ...interface{}) {
	defaultLogger.cLog(ctx, 1, INFO, args)
}

func CInfoF(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.cLogF(ctx, 1, INFO, format, args)
}

func Warn(args ...interface{}) {
	defaultLogger.log(1, WARN, args)
}

func WarnF(format string, args ...interface{}) {
	defaultLogger.logF(1, WARN, format, args)
}

func CWarn(ctx context.Context, args ...interface{}) {
	defaultLogger.cLog(ctx, 1, WARN, args)
}

func CWarnF(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.cLogF(ctx, 1, WARN, format, args)
}

func Error(args ...interface{}) {
	defaultLogger.log(1, ERROR, args)
}

func ErrorF(format string, args ...interface{}) {
	defaultLogger.logF(1, ERROR, format, args)
}

func CError(ctx context.Context, args ...interface{}) {
	defaultLogger.cLog(ctx, 1, ERROR, args)
}

func CErrorF(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.cLogF(ctx, 1, ERROR, format, args)
}

func Fatal(args ...interface{}) {
	defaultLogger.log(1, FATAL, args)
}

func FatalF(format string, args ...interface{}) {
	defaultLogger.logF(1, FATAL, format, args)
}

func CFatal(ctx context.Context, args ...interface{}) {
	defaultLogger.cLog(ctx, 1, FATAL, args)
}

func CFatalF(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.cLogF(ctx, 1, FATAL, format, args)
}

func SetConsoleLevel(level Level) {
	defaultLogger.config.consoleLevel = level
}

func SetFileLevel(level Level) {
	defaultLogger.SetFileLevel(level)
}

func SetFilePath(path string) {
	defaultLogger.SetFilePath(path)
	fmt.Println("change log file", defaultLogger.config.GetFilePath())
}

func SetFileSize(size int64) {
	defaultLogger.SetFileSize(size)
}

func SetFileCount(count int) {
	defaultLogger.SetFileCount(count)
}

func SetContextCallBackFunc(f func(ctx context.Context) string) {
	defaultLogger.SetContextCallBackFunc(f)
}

func SetCallBackFunc(f func(*Format)) {
	defaultLogger.SetCallBackFunc(f)
}

func SetConsoleFormat(f func(*Format) string) {
	defaultLogger.SetConsoleFormat(f)
}

func SetFileFormat(f func(*Format) string) {
	defaultLogger.SetFileFormat(f)
}

func SetFileBaseName(name string) {
	defaultLogger.SetFileBaseName(name)
}
