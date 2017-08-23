package logger

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type Format struct {
	Time    time.Time
	Level   Level
	Message string
	File    string
	Line    int
}

func NewFormat(level Level, v []interface{}, skip int) *Format {
	var format Format
	format.Time = time.Now()
	format.Level = level

	for _, _v := range v {
		format.Message += fmt.Sprint(_v) + " "
	}
	format.Message = format.Message[:len(format.Message)-1]

	_, format.File, format.Line, _ = runtime.Caller(skip)
	short := format.File
	for i := len(format.File) - 1; i > 0; i-- {
		if format.File[i] == '/' {
			short = format.File[i+1:]
			break
		}
	}
	format.File = short

	return &format
}

func (this *Format) String() string {
	return this.Time.Format("2006-01-02 15:04:05") + " [" + this.Level.String() + "] " + this.Message + " --" + this.File + ":" + strconv.Itoa(this.Line)
}

func (this *Format) ColorString() string {
	cn := "37"
	switch this.Level {
	case DEBUG:
		cn = "34"
	case INFO:
		cn = "32"
	case WARN:
		cn = "33"
	case ERROR:
		cn = "35"
	case FATAL:
		cn = "31"
	}
	return this.Time.Format("2006-01-02 15:04:05") + " [\033[;" + cn + "m" + this.Level.String() + "\033[0m] " + this.Message + " --" + this.File + ":" + strconv.Itoa(this.Line)
}
