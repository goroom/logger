package logger

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Logger struct {
	sync.Mutex
	config      *Config
	file        *os.File
	fileName    string
	fileSize    int64
	maxFileSize int64
	formatFunc  func(*Format) string
}

func NewLogger(config *Config) (*Logger, error) {
	if config == nil {
		config = NewDefaultConfig()
	}
	var logger Logger
	logger.config = config
	logger.fileName = logger.config.FilePath + "/" + logger.config.FileBaseName
	logger.maxFileSize = logger.config.MaxFileSize * int64(logger.config.MaxFileSizeUnit)
	//	fmt.Println("Max file size", logger.maxFileSize)

	if logger.config.Level != OFF {
		err := os.MkdirAll(logger.config.FilePath, 0777)
		if err != nil {
			return nil, err
		}
		logger.file, err = os.OpenFile(logger.fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return nil, err
		}

		logger.fileSize = logger.getFileSize(logger.fileName)
		//		fmt.Println("File size", logger.fileSize)
	}

	return &logger, nil
}

func (this *Logger) getFileSize(file_path string) int64 {
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
		if this.IsFileExit(file_path) {
			os.Remove(file_path)
		}

		for i := this.config.MaxFileCount; i > 0; i-- {
			file_path := this.fileName + "." + strconv.Itoa(i-2)
			if this.IsFileExit(file_path) {
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

func (this *Logger) IsFileExit(file_path string) bool {
	_, err := os.Stat(file_path)
	return err == nil || os.IsExist(err)
}

func (this *Logger) SetFormatFunc(f func(*Format) string) {
	this.formatFunc = f
}

func (this *Logger) Log(skip int, level Level, v []interface{}) {
	this.Lock()
	defer this.Unlock()
	if this.config.Level > level && this.config.ConsoleLevel > level && this.config.CallBackFunc == nil {
		return
	}

	format := NewFormat(level, v, skip+2)

	if this.config.Level <= level || this.config.ConsoleLevel <= level {
		format_string := ""

		if this.formatFunc != nil {
			format_string = this.formatFunc(format)
		} else {
			format_string = format.ColorString() + "\n"
		}

		if this.config.ConsoleLevel <= level {
			fmt.Print(format_string)
		}

		if this.config.Level <= level {
			this.Write(&format_string)
		}
	}

	if this.config.CallBackFunc != nil {
		this.config.CallBackFunc(format)
	}
}

func (this *Logger) Write(format_string *string) {
	err := this.checkFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	this.file.Write([]byte(*format_string))

	this.fileSize += int64(len(*format_string))
}

func (this *Logger) Debug(v ...interface{}) {
	this.Log(1, DEBUG, v)
}

func (this *Logger) Info(v ...interface{}) {
	this.Log(1, INFO, v)
}

func (this *Logger) Warn(v ...interface{}) {
	this.Log(1, WARN, v)
}
func (this *Logger) Error(v ...interface{}) {
	this.Log(1, ERROR, v)
}

func (this *Logger) Fatal(v ...interface{}) {
	this.Log(1, FATAL, v)
}

func (this *Logger) SetLevel(level Level) {
	this.Lock()
	defer this.Unlock()

	this.config.Level = level
}

func (this *Logger) SetConsoleLevel(level Level) {
	this.Lock()
	defer this.Unlock()

	this.config.ConsoleLevel = level
}

func (this *Logger) SetMaxFileSize(size int64, unit Unit) {
	this.Lock()
	defer this.Unlock()

	this.config.MaxFileSize = size
	this.config.MaxFileSizeUnit = unit
	this.maxFileSize = this.config.MaxFileSize * int64(this.config.MaxFileSizeUnit)
}

func (this *Logger) SetCallBackFunc(f func(*Format)) {
	this.Lock()
	defer this.Unlock()

	this.config.CallBackFunc = f
}
