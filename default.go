package logger

var defaultLogger *Logger

func init() {
	defaultLogger = NewLogger()
}

func GetDefaultLogger() *Logger {
	return defaultLogger
}

func InitDefaultLogger(opts ...Options) {
	defaultLogger.Close()
	defaultLogger = NewLogger(opts...)
}

func Debug(args ...interface{}) {
	defaultLogger.log(1, DEBUG, args)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.logf(1, DEBUG, format, args)
}

func Info(args ...interface{}) {
	defaultLogger.log(1, INFO, args)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.logf(1, INFO, format, args)
}

func Warn(args ...interface{}) {
	defaultLogger.log(1, WARN, args)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.logf(1, WARN, format, args)
}

func Error(args ...interface{}) {
	defaultLogger.log(1, ERROR, args)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.logf(1, ERROR, format, args)
}

func Wait() {
	defaultLogger.Wait()
}
