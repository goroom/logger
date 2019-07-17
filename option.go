package logger

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type option struct {
	ChannelBuffLength int64 // 异步缓存大小

	FileName       string
	FileLevel      Level
	FileDir        string
	FileSplit      func(*Logger, *File, *Format) (*File, error)
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

		FileName:       strings.TrimRight(path.Base(filepath.ToSlash(os.Args[0])), ".exe"),
		FileLevel:      OFF,
		FileDir:        "./log",
		FileSplit:      DefaultFileSplitBySize,
		FileSizeMax:    MB * 50,
		FileCountMax:   10,
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
	if !isShowConsole {
		return func(o *option) { o.ConsoleLevel = OFF }
	}
	return func(o *option) {}
}

// WithFileDir 设置文件路径
func WithFileDir(dir string) Options {
	return func(o *option) { o.FileDir = dir }
}

// WithFileName 设置文件名
func WithFileName(fileName string) Options {
	return func(o *option) { o.FileName = fileName }
}

// WithFileSize 设置文件大小
func WithFileSize(fileSize Unit) Options {
	return func(o *option) { o.FileSizeMax = fileSize }
}

// WithFileSize 设置最多保留的文件个数
func WithFileMaxCount(maxCount int) Options {
	return func(o *option) { o.FileCountMax = maxCount }
}

// WithFileSize 设置自定义的文件拆分方法
// 默认按照文件大小拆分
func WithFileSplit(f func(*Logger, *File, *Format) (*File, error)) Options {
	return func(o *option) { o.FileSplit = f }
}

// WithCallBack 回调
func WithCallBack(f func(*Format)) Options {
	return func(o *option) { o.CallBackFunc = f }
}
