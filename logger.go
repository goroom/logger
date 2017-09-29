package logger

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Logger struct {
	sync.Mutex
	isInit           bool
	config           *Config
	file             *os.File
	fileName         string
	fileSize         int64
	maxFileSize      int64
	consoleFormatFun func(*Format) string
	fileFormatFun    func(*Format) string
	ch               chan *Format
}

var (
	gLogger *Logger
)

func (l *Logger) init(config *Config) error {
	if config == nil {
		return fmt.Errorf("Config is nil")
	}
	if l.isInit {
		return fmt.Errorf("Repeat init logger")
	}
	l.config = config

	l.fileName = l.config.FilePath + "/" + l.config.FileBaseName
	l.maxFileSize = l.config.MaxFileSize * int64(l.config.MaxFileSizeUnit)
	l.ch = make(chan *Format, 10000)

	if l.config.FileLevel != OFF {
		err := os.MkdirAll(l.config.FilePath, 0777)
		if err != nil {
			return err
		}
		l.file, err = os.OpenFile(l.fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return err
		}

		l.fileSize = getFileSize(l.fileName)
	}

	go func() {
		var format *Format
		for {
			select {
			case format = <-gLogger.ch:
				var formatString string
				if gLogger.fileFormatFun != nil {
					formatString = gLogger.fileFormatFun(format)
				} else {
					formatString = format.FileString()
				}
				formatString += "\n"
				gLogger.write(&formatString)
			}
		}
	}()

	l.isInit = true
	return nil
}

func checkDefaultLogger() {
	if gLogger == nil {
		err := Init(nil)
		if err != nil {
			fmt.Println("初始化日志出错", err)
		}
	}
}

func Init(config *Config) error {
	if config == nil {
		config = NewDefaultConfig()
	}
	gLogger = &Logger{}
	return gLogger.init(config)
}

func getFileSize(file_path string) int64 {
	file_info, err := os.Stat(file_path)
	if err != nil {
		return 0
	}
	return file_info.Size()
}

func (this *Logger) checkFile() error {
	if this.fileSize >= this.maxFileSize {
		this.file.Close()
		file_path := this.fileName + "." + strconv.Itoa(this.config.MaxFileCount-1)
		if this.isFileExit(file_path) {
			os.Remove(file_path)
		}

		for i := this.config.MaxFileCount; i > 0; i-- {
			file_path := this.fileName + "." + strconv.Itoa(i-2)
			if this.isFileExit(file_path) {
				os.Rename(file_path, this.fileName+"."+strconv.Itoa(i-1))
				//				fmt.Println(file_path, "=>", this.fileName+"."+strconv.Itoa(i-1))
			}
		}

		os.Rename(this.fileName, this.fileName+".1")
		os.Remove(this.fileName)
		var err error
		this.file, err = os.OpenFile(this.fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		this.fileSize = 0
	}
	return nil
}

func (this *Logger) isFileExit(file_path string) bool {
	_, err := os.Stat(file_path)
	return err == nil || os.IsExist(err)
}

func (this *Logger) SetFormatFunc(f func(*Format) string) {
	this.fileFormatFun = f
	this.consoleFormatFun = f
}

func log(skip int, level Level, v []interface{}) {
	checkDefaultLogger()
	var format *Format

	//Console log
	if level >= gLogger.config.ConsoleLevel {

		format = NewFormat(level, v, skip+2)

		var formatString string
		if gLogger.consoleFormatFun != nil {
			formatString = gLogger.consoleFormatFun(format)
		} else {
			formatString = format.ConsoleString()
		}
		fmt.Println(formatString)
	}

	//File log
	if level >= gLogger.config.FileLevel {
		if format == nil {
			format = NewFormat(level, v, skip+2)
		}
		gLogger.ch <- format
	}

	//Call back
	if gLogger.config.CallBackFunc != nil {
		if format == nil {
			format = NewFormat(level, v, skip+2)
		}
		gLogger.config.CallBackFunc(format)
	}
}

func (this *Logger) write(format_string *string) {
	err := this.checkFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	this.file.Write([]byte(*format_string))

	this.fileSize += int64(len(*format_string))
}

func Debug(v ...interface{}) {
	log(1, DEBUG, v)
}

func Debugf(format string, a ...interface{}) {
	log(1, DEBUG, []interface{}{fmt.Sprintf(format, a...)})
}

func Info(v ...interface{}) {
	log(1, INFO, v)
}

func Infof(format string, a ...interface{}) {
	log(1, INFO, []interface{}{fmt.Sprintf(format, a...)})
}

func Warn(v ...interface{}) {
	log(1, WARN, v)
}

func Warnf(format string, a ...interface{}) {
	log(1, WARN, []interface{}{fmt.Sprintf(format, a...)})
}

func Error(v ...interface{}) {
	log(1, ERROR, v)
}

func Errorf(format string, a ...interface{}) {
	log(1, ERROR, []interface{}{fmt.Sprintf(format, a...)})
}

func Fatal(v ...interface{}) {
	log(1, FATAL, v)
}

func Fatalf(format string, a ...interface{}) {
	log(1, FATAL, []interface{}{fmt.Sprintf(format, a...)})
}

func SetFileLevel(level Level) {
	gLogger.Lock()
	defer gLogger.Unlock()

	gLogger.config.FileLevel = level
}

func SetConsoleLevel(level Level) {
	gLogger.Lock()
	defer gLogger.Unlock()

	gLogger.config.ConsoleLevel = level
}

func SetMaxFileSize(size int64, unit Unit) {
	gLogger.Lock()
	defer gLogger.Unlock()

	gLogger.config.MaxFileSize = size
	gLogger.config.MaxFileSizeUnit = unit
	gLogger.maxFileSize = gLogger.config.MaxFileSize * int64(gLogger.config.MaxFileSizeUnit)
}

func SetCallBackFunc(f func(*Format)) {
	gLogger.Lock()
	defer gLogger.Unlock()

	gLogger.config.CallBackFunc = f
}
