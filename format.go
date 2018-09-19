package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type Format struct {
	Time  time.Time
	Level Level
	Args  []interface{}
	File  string
	Line  int
}

func NewFormat(level Level, args []interface{}, skip int) *Format {
	format := Format{
		Time:  time.Now(),
		Level: level,
		Args:  args,
	}

	_, format.File, format.Line, _ = runtime.Caller(skip)

	return &format
}

func (f *Format) ArgsDefaultFormat() string {
	var s []string
	for i := 0; i < len(f.Args); i++ {
		s = append(s, fmt.Sprint(f.Args[i]))
	}
	return strings.Join(s, " ")
}

type FormatFunc func(*Format) string

func defaultFormatFileName(fName string) string {
	return path.Base(fName)
}

func defaultConsoleFormatFunc(f *Format) string {
	return fmt.Sprintf("%s \033[;%dm%s\033[0m %s -%s:%d",
		f.Time.Format("2006-01-02 15:04:05"), f.Level.ConsoleColorNum(), f.Level.String(),
		f.ArgsDefaultFormat(), defaultFormatFileName(f.File), f.Line)
}

func defaultFileFormatFunc(f *Format) string {
	return fmt.Sprintf("%s %s %s -%s:%d\n",
		f.Time.Format("2006-01-02 15:04:05"), f.Level.String(),
		f.ArgsDefaultFormat(), defaultFormatFileName(f.File), f.Line)
}
