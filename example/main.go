package main

import (
	"context"
	"fmt"
	"github.com/goroom/logger"
	"time"
)

func main() {
	logger.Debug("no save file")

	logger.SetFileLevel(logger.ALL)       // default off
	logger.SetFileSize(logger.KB.CalB(2)) // default 5MB
	logger.SetFileCount(3)                // default 5
	logger.Debug("d")
	logger.DebugF("%s %d", "a", 1)
	logger.Info("i")
	logger.InfoF("%s %d", "a", 1)
	logger.Warn("w")
	logger.WarnF("%s %d", "a", 1)
	logger.Error("e")
	logger.ErrorF("%s %d", "a", 1)
	logger.Fatal("f")
	logger.FatalF("%s %d", "a", 1)

	logger.SetContextCallBackFunc(func(ct context.Context) string {
		return "id"
	})
	logger.CDebug(context.Background(), "d")
	logger.CDebugF(context.Background(), "%s %d", "a", 1)
	logger.CInfo(context.Background(), "i")
	logger.CInfoF(context.Background(), "%s %d", "a", 1)
	logger.CWarn(context.Background(), "w")
	logger.CWarnF(context.Background(), "%s %d", "a", 1)
	logger.CError(context.Background(), "e")
	logger.CErrorF(context.Background(), "%s %d", "a", 1)
	logger.CFatal(context.Background(), "f")
	logger.CFatalF(context.Background(), "%s %d", "a", 1)

	logger.SetConsoleFormat(func(f *logger.Format) string {
		return fmt.Sprintf("%s %v", f.Time.Format("15:02:03"), f.ArgsDefaultFormat())
	})
	logger.Debug("customer console format")

	logger.SetCallBackFunc(func(f *logger.Format) {
		fmt.Println(f)
	})
	logger.Debug("has call back")

	customerLogger := logger.NewLogger("tag1")
	customerLogger.SetFileLevel(logger.ALL)
	customerLogger.SetFilePath("./log2")
	customerLogger.SetFileBaseName("ex")
	customerLogger.Debug(1)
	time.Sleep(1e9)
}
