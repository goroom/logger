package logger

import (
	"context"
	"os"
	"path/filepath"
)

type config struct {
	consoleLevel      Level
	consoleFormatFunc FormatFunc

	fileLevel      Level
	fileFormatFunc FormatFunc
	fileChanCnt    int32
	filePath       string
	fileNameBase   string
	fileSizeMax    int64
	fileCntMax     int

	ctxCBFunc func(context.Context) string
}

func (c *config) GetFilePath() string {
	return c.filePath + "/" + c.fileNameBase + ".log"
}

func getDefaultConfig() *config {
	execFilePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil || execFilePath == "" {
		execFilePath = "."
	}
	execFileName := filepath.Base(os.Args[0])
	if execFileName == "" {
		execFileName = "unknown"
	}
	return &config{
		consoleLevel:      ALL,
		consoleFormatFunc: defaultConsoleFormatFunc,

		fileLevel:      OFF,
		fileFormatFunc: defaultFileFormatFunc,
		fileChanCnt:    1000000,
		filePath:       execFilePath + "/log",
		fileNameBase:   execFileName,
		fileSizeMax:    MB.CalB(5),
		fileCntMax:     5,
	}
}
