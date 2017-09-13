package main

import (
	"fmt"
	"time"

	"github.com/goroom/logger"
)

type A struct {
	a int
	b string
}

func main() {
	//获取一个默认配置
	default_config := logger.NewDefaultConfig()
	//修改默认配置的日志文件路径
	default_config.FilePath = "./log"
	//修改默认配置的日志文件名称
	default_config.FileBaseName = "logger.log"
	//设置日志级别的两种方式
	default_config.ConsoleLevel = logger.StringLevel("ALL")
	default_config.FileLevel = logger.WARN

	//初始化日志
	err := logger.Init(default_config)
	if err != nil {
		fmt.Println(err)
		return
	}
	//修改日志文件最大为1KB
	logger.SetMaxFileSize(1, logger.KB)

	logger.Debug("这条日志会出现在终端，不会出现在日志文件。")
	logger.Warn("这条日志会出现在终端和日志文件。")

	//输出debug日志
	logger.Debug("debug_message", A{a: 1, b: "2"})

	//设置文件日志级别为INFO
	logger.SetFileLevel(logger.INFO)
	//设置终端输出的日志级别为INFO
	logger.SetConsoleLevel(logger.INFO)
	//打印日志
	logger.Debug("debug_message")
	logger.Info("info_message")
	logger.Warn("warn_message")

	//设置日志回调，可以自行处理日志，例如发送到日志服务器。
	logger.SetCallBackFunc(func(f *logger.Format) {
		fmt.Println("[CB]", f.ConsoleString())
	})

	//打印日志
	logger.Error("error_message")
	logger.Fatal("fatal_message")

	time.Sleep(1e9)
}
