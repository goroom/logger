package logger

import (
	"bytes"
	"encoding/json"
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
	_, file, line, _ := runtime.Caller(skip)

	return &Format{
		Time:  time.Now(),
		Level: level,
		Args:  args,
		File:  file,
		Line:  line,
	}
}

func (f *Format) ArgsDefaultFormat() []byte {
	var buffer bytes.Buffer
	for i := 0; i < len(f.Args); i++ {
		s := fmt.Sprint(f.Args[i])
		if len(s) == 12 && strings.HasPrefix(s, "0x") {
			data, err := json.Marshal(f.Args[i])
			if err == nil {
				s = "&" + string(data)
			}
		}
		buffer.WriteString(s + " ")
	}
	return buffer.Bytes()
}

func defaultFormatFileName(fName string) string {
	return path.Base(fName)
}

func defaultConsoleFormatFunc(f *Format) string {
	return fmt.Sprintf("%s [%s] %s-%s:%d", f.Time.Format("2006-01-02 15:04:05"), f.Level.ConsoleColorString(),
		f.ArgsDefaultFormat(), defaultFormatFileName(f.File), f.Line)
}

func defaultFileFormatFunc(f *Format) []byte {
	return []byte(fmt.Sprintf("%s [%s] %s-%s:%d\n", f.Time.Format("2006-01-02 15:04:05"), f.Level.String(),
		f.ArgsDefaultFormat(), defaultFormatFileName(f.File), f.Line))
}
