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
