package logger

import (
	"fmt"
	"golang.org/x/net/context"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
)

var mapLogger = make(map[string]*Logger)

func GetLogger(tag string) *Logger {
	l, ok := mapLogger[tag]
	if !ok {
		return nil
	}
	return l
}

func NewLogger(tag string) *Logger {
	l := GetLogger(tag)
	if l != nil {
		return nil
	}
	l = newLogger(tag)
	mapLogger[tag] = l
	return l
}

func newLogger(tag string) *Logger {
	defaultConfig := getDefaultConfig()
	logger := Logger{
		tag:    tag,
		config: defaultConfig,

		fileFormatChan: make(chan *Format, defaultConfig.fileChanCnt),
	}
	logger.run()
	return &logger
}

type Logger struct {
	tag    string
	config *config

	fileFormatChanCurrCnt int32
	fileFormatChan        chan *Format
	currFileSize          int64
	file                  *os.File

	callBackFunc func(*Format)
}

func (l *Logger) log(skip int, level Level, args []interface{}) {
	if l == nil {
		return
	}
	if level < l.config.consoleLevel && level < l.config.fileLevel {
		return
	}

	format := NewFormat(level, args, skip+2)

	if level >= l.config.consoleLevel {
		fmt.Println(l.config.consoleFormatFunc(format))
	}

	if level >= l.config.fileLevel {
		if l.fileFormatChanCurrCnt < l.config.fileChanCnt {
			atomic.AddInt32(&l.fileFormatChanCurrCnt, 1)
			l.fileFormatChan <- format
		} else {
			fmt.Println("Log stack overflow, discard.")
		}
	}

	if l.callBackFunc != nil {
		l.callBackFunc(format)
	}
}

func (l *Logger) cLog(ctx context.Context, skip int, level Level, args []interface{}) {
	args = append([]interface{}{l.getLogID(ctx)}, args...)
	l.log(skip+1, level, args)
}

func (l *Logger) logF(skip int, level Level, format string, args []interface{}) {
	l.log(skip+1, level, []interface{}{fmt.Sprintf(format, args...)})
}

func (l *Logger) cLogF(ctx context.Context, skip int, level Level, format string, args []interface{}) {
	l.cLog(ctx, skip+1, level, []interface{}{fmt.Sprintf(format, args...)})
}

func (l *Logger) Debug(args ...interface{}) {
	l.log(1, DEBUG, args)
}

func (l *Logger) DebugF(format string, args ...interface{}) {
	l.logF(1, DEBUG, format, args)
}

func (l *Logger) CDebug(ctx context.Context, args ...interface{}) {
	l.cLog(ctx, 1, DEBUG, args)
}

func (l *Logger) CDebugF(ctx context.Context, format string, args ...interface{}) {
	l.cLogF(ctx, 1, DEBUG, format, args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.log(1, WARN, args)
}

func (l *Logger) WarnF(format string, args ...interface{}) {
	l.logF(1, WARN, format, args)
}

func (l *Logger) CWarn(ctx context.Context, args ...interface{}) {
	l.cLog(ctx, 1, WARN, args)
}

func (l *Logger) CWarnF(ctx context.Context, format string, args ...interface{}) {
	l.cLogF(ctx, 1, WARN, format, args)
}

func (l *Logger) Error(args ...interface{}) {
	l.log(1, ERROR, args)
}

func (l *Logger) ErrorF(format string, args ...interface{}) {
	l.logF(1, ERROR, format, args)
}

func (l *Logger) CError(ctx context.Context, args ...interface{}) {
	l.cLog(ctx, 1, ERROR, args)
}

func (l *Logger) CErrorF(ctx context.Context, format string, args ...interface{}) {
	l.cLogF(ctx, 1, ERROR, format, args)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.log(1, FATAL, args)
}

func (l *Logger) FatalF(format string, args ...interface{}) {
	l.logF(1, FATAL, format, args)
}

func (l *Logger) CFatal(ctx context.Context, args ...interface{}) {
	l.cLog(ctx, 1, FATAL, args)
}

func (l *Logger) CFatalF(ctx context.Context, format string, args ...interface{}) {
	l.cLogF(ctx, 1, FATAL, format, args)
}

func (l *Logger) Info(args ...interface{}) {
	l.log(1, INFO, args)
}

func (l *Logger) CInfo(ctx context.Context, args ...interface{}) {
	l.cLog(ctx, 1, INFO, args)
}

func (l *Logger) getLogID(ctx context.Context) string {
	if l.config.ctxCBFunc == nil || ctx == nil {
		return "-"
	}
	logID := l.config.ctxCBFunc(ctx)
	if logID == "" {
		logID = "-"
	}
	return logID
}

func (l *Logger) SetFileLevel(level Level) {
	l.config.fileLevel = level
}

func (l *Logger) SetFilePath(path string) error {
	filePath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	l.config.filePath = filePath
	return nil
}

func (l *Logger) SetFileBaseName(name string) error {
	if strings.Index(name, "/") >= 0 || strings.Index(name, "\\") >= 0 {
		return fmt.Errorf("illegal file name")
	}
	l.config.fileNameBase = name
	return nil
}

func (l *Logger) run() {
	go func() {
		var format *Format
		var ok bool
		for {
			select {
			case format, ok = <-l.fileFormatChan:
				if !ok {
					return
				}
				atomic.AddInt32(&l.fileFormatChanCurrCnt, -1)
				formatString := l.config.fileFormatFunc(format)
				l.writeFile([]byte(formatString))
			}
		}
	}()
}

func (l *Logger) writeFile(data []byte) {
	if err := l.checkFile(); err != nil {
		fmt.Println(err)
		return
	}
	l.file.Write(data)
	atomic.AddInt64(&l.currFileSize, int64(len(data)))
}

func (l *Logger) openFile() error {
	var err error
	err = os.MkdirAll(l.config.filePath, 0777)
	if err != nil {
		return err
	}
	l.file, err = os.OpenFile(l.config.GetFilePath(), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	l.currFileSize = 0
	return nil
}

func (l *Logger) checkFile() error {
	if l.file == nil {
		return l.openFile()
	}
	if l.currFileSize < l.config.fileSizeMax {
		return nil
	}
	err := l.file.Close()
	if err != nil {
		return err
	}
	logFilePath := l.config.GetFilePath()
	filePath := logFilePath + "." + strconv.Itoa(l.config.fileCntMax-1)
	if isFileExit(filePath) {
		// fmt.Println("Remove oldest log file", filePath)
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	for i := l.config.fileCntMax; i > 0; i-- {
		tFilePath := logFilePath + "." + strconv.Itoa(i-2)
		if isFileExit(tFilePath) {
			// fmt.Println("Rename log file", tFilePath, logFilePath+"."+strconv.Itoa(i-1))
			err := os.Rename(tFilePath, logFilePath+"."+strconv.Itoa(i-1))
			if err != nil {
				return err
			}
		}
	}

	// fmt.Println("Rename curr log file", logFilePath, logFilePath+".1")
	err = os.Rename(logFilePath, logFilePath+".1")
	if err != nil {
		return err
	}
	return l.openFile()
}

func (l *Logger) SetFileSize(size int64) {
	l.config.fileSizeMax = size
}

func (l *Logger) SetFileCount(count int) {
	l.config.fileCntMax = count
}

func (l *Logger) SetContextCallBackFunc(f func(ctx context.Context) string) {
	if l == nil {
		return
	}
	l.config.ctxCBFunc = f
}

func (l *Logger) SetCallBackFunc(f func(*Format)) {
	l.callBackFunc = f
}
