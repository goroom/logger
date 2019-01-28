package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var bufferPool *sync.Pool

func init() {
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

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

type FormatFunc func(*Format) []byte

func defaultFormatFileName(fName string) string {
	return path.Base(fName)
}

func defaultConsoleFormatFunc(f *Format) []byte {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer bufferPool.Put(buffer)
	buffer.WriteString(f.Time.Format("2006-01-02 15:04:05") + " " + f.Level.ConsoleColorString() + " ")
	buffer.Write(f.ArgsDefaultFormat())
	buffer.WriteString("-" + defaultFormatFileName(f.File) + ":" + strconv.Itoa(f.Line))
	return buffer.Bytes()
}

func defaultFileFormatFunc(f *Format) []byte {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer bufferPool.Put(buffer)
	buffer.WriteString(f.Time.Format("2006-01-02 15:04:05") + " " + f.Level.String() + " ")
	buffer.Write(f.ArgsDefaultFormat())
	buffer.WriteString("-" + defaultFormatFileName(f.File) + ":" + strconv.Itoa(f.Line) + "\n")
	return buffer.Bytes()
}
