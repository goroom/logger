package logger

import (
	"fmt"
	"os"
	"path"
)

type option struct {
	ChannelBuffLength int64 // 异步缓存大小

	FileName       string
	FileLevel      Level
	FileDir        string
	FileSplit      func(*Logger, *File) (*File, error)
	FileSizeMax    Unit
	FileCountMax   int
	FileFormatFunc func(*Format) []byte

	ConsoleLevel      Level
	ConsoleFormatFunc func(*Format) string

	CallBackFunc func(*Format)
}

type Options func(*option)

func newDefaultOption() *option {
	return &option{
		ChannelBuffLength: 10000,

		FileName:       path.Base(os.Args[0]),
		FileLevel:      OFF,
		FileDir:        "./log",
		FileSplit:      DefaultFileSplitBySize,
		FileSizeMax:    MB * 100,
		FileCountMax:   5,
		FileFormatFunc: defaultFileFormatFunc,

		ConsoleLevel:      INFO,
		ConsoleFormatFunc: defaultConsoleFormatFunc,
	}
}

func (o *option) GetFilePath(no int) string {
	if no == 0 {
		return fmt.Sprintf("%s/%s.log", o.FileDir, o.FileName)
	}
	return fmt.Sprintf("%s/%s_%d.log", o.FileDir, o.FileName, no)
}

// WithFileLevel 设置文件输出级别
func WithFileLevel(lvl Level) Options {
	return func(o *option) { o.FileLevel = lvl }
}

// WithConsoleLevel 设置终端输出级别
func WithConsoleLevel(lvl Level) Options {
	return func(o *option) { o.ConsoleLevel = lvl }
}

// WithConsole 是否在终端输出日志
func WithConsole(isShowConsole bool) Options {
	return func(o *option) { o.ConsoleLevel = OFF }
}

// WithFileDir 设置文件路径
func WithFileDir(dir string) Options {
	return func(o *option) { o.FileDir = dir }
}

func WithFileName(fileName string) Options {
	return func(o *option) { o.FileName = fileName }
}

// WithFileSize 设置文件大小
func WithFileSize(fileSize Unit) Options {
	return func(o *option) { o.FileSizeMax = fileSize }
}

// WithFileSize 设置文件大小
func WithFileSplit(f func(*Logger, *File) (*File, error)) Options {
	return func(o *option) { o.FileSplit = f }
}

// WithCallBack 回调
func WithCallBack(f func(*Format)) Options {
	return func(o *option) { o.CallBackFunc = f }
}
