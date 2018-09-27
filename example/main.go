package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime/trace"
	"time"

	"github.com/goroom/logger"
)

func test() {
	f, _ := os.Create("trace.out")
	trace.Start(f)

	logger.SetConsoleLevel(logger.OFF)
	logger.SetFileLevel(logger.ALL)
	logger.SetFileCount(5)
	t1 := time.Now().UnixNano()
	go func() {
		for {
			time.Sleep(1e9)
			fmt.Println(logger.GetDefaultLogger().GetFileChannelCount())
		}
	}()
	go func() {
		for {
			time.Sleep(1e5)
			if logger.GetDefaultLogger().GetFileChannelCount() <= 0 {
				fmt.Println((time.Now().UnixNano() - t1) / 1e6)
				os.Exit(0)
			}
		}
	}()
	for i := 0; i < 1000000; i++ {
		//time.Sleep(1e3)
		//logger.Debug("12ks0192j192hisj12hs1029")
		logger.Debug("logger.GetDefaultLogger()")
	}
	trace.Stop()
	select {}
}

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

	logger.SetConsoleFormat(func(f *logger.Format) []byte {
		return []byte(fmt.Sprintf("%s %s", f.Time.Format("15:02:03"), string(f.ArgsDefaultFormat())))
	})
	logger.Debug("customer console format")

	logger.SetCallBackFunc(func(f *logger.Format) {
		b, _ := json.Marshal(f)
		fmt.Println(string(b))
	})
	logger.Debug("has call back")

	customerLogger := logger.NewLogger("tag1")
	customerLogger.SetFileLevel(logger.ALL)
	customerLogger.SetFilePath("./log2")
	customerLogger.SetFileBaseName("ex")
	customerLogger.Debug("customer logger")
	for logger.GetDefaultLogger().GetFileChannelCount() > 0 {
		time.Sleep(1e5)
	}

	for customerLogger.GetFileChannelCount() > 0 {
		time.Sleep(1e5)
	}
}
