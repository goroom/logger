package logger

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Logger struct {
	opt *option

	fileFormatChanCurrCnt int64
	fileFormatChan        chan *Format

	file *File
}

func NewLogger(opts ...Options) *Logger {
	opt := newDefaultOption()
	for _, v := range opts {
		v(opt)
	}

	lg := &Logger{
		opt:            opt,
		fileFormatChan: make(chan *Format, opt.ChannelBuffLength),
	}
	lg.run()

	return lg
}

func (l *Logger) Close() {
	if l == nil {
		return
	}
	close(l.fileFormatChan)
	l.fileFormatChanCurrCnt = 0
	l.file.Close()
}

func (l *Logger) log(skip int, level Level, args []interface{}) {
	if l == nil {
		return
	}

	format := NewFormat(level, args, skip+2)

	if level >= l.opt.ConsoleLevel {
		fmt.Println(l.opt.ConsoleFormatFunc(format))
	}

	if level >= l.opt.FileLevel {
		if atomic.LoadInt64(&l.fileFormatChanCurrCnt) < l.opt.ChannelBuffLength {
			atomic.AddInt64(&l.fileFormatChanCurrCnt, 1)
			l.fileFormatChan <- format
		} else {
			fmt.Println("Log stack overflow, discard.")
		}
	}

	if l.opt.CallBackFunc != nil {
		l.opt.CallBackFunc(format)
	}
}

func (l *Logger) Wait() {
	for {
		if l.fileFormatChanCurrCnt == 0 {
			return
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func (l *Logger) isLowLevel(lvl Level) bool {
	if lvl < l.opt.FileLevel && lvl < l.opt.ConsoleLevel && l.opt.CallBackFunc == nil {
		return true
	}
	return false
}

func (l *Logger) logf(skip int, level Level, format string, args []interface{}) {
	l.log(skip+1, level, []interface{}{fmt.Sprintf(format, args...)})
}

func (l *Logger) Debug(args ...interface{}) {
	if l.isLowLevel(DEBUG) {
		return
	}
	l.log(1, DEBUG, args)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.isLowLevel(DEBUG) {
		return
	}
	l.logf(1, DEBUG, format, args)
}

func (l *Logger) Info(args ...interface{}) {
	if l.isLowLevel(INFO) {
		return
	}
	l.log(1, INFO, args)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.isLowLevel(INFO) {
		return
	}
	l.logf(1, INFO, format, args)
}

func (l *Logger) Warn(args ...interface{}) {
	if l.isLowLevel(WARN) {
		return
	}
	l.log(1, WARN, args)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.isLowLevel(WARN) {
		return
	}
	l.logf(1, WARN, format, args)
}

func (l *Logger) Error(args ...interface{}) {
	if l.isLowLevel(ERROR) {
		return
	}
	l.log(1, ERROR, args)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.isLowLevel(ERROR) {
		return
	}
	l.logf(1, ERROR, format, args)
}

func (l *Logger) SetFileLevel(level Level) {
	l.opt.FileLevel = level
}

func (l *Logger) SetConsoleLevel(level Level) {
	l.opt.ConsoleLevel = level
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
				l.writeFile(l.opt.FileFormatFunc(format))
				atomic.AddInt64(&l.fileFormatChanCurrCnt, -1)
			}
		}
	}()
}

func (l *Logger) writeFile(data []byte) {
	f, err := l.opt.FileSplit(l, l.file)
	if err != nil {
		fmt.Println(err)
		return
	}
	l.file = f

	err = l.file.write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (l *Logger) SetFileSize(size Unit) {
	l.opt.FileSizeMax = size
}

func (l *Logger) SetFileCount(count int) {
	l.opt.FileCountMax = count
}
