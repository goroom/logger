package main

import (
	"fmt"

	"github.com/lesterpang/logger"
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
	default_config.Level = logger.StringLevel("DEBUG")
	default_config.Level = logger.DEBUG

	//新建日志
	log, err := logger.NewLogger(default_config)
	if err != nil {
		fmt.Println(err)
		return
	}
	//修改日志文件最大为1KB
	log.SetMaxFileSize(1, logger.KB)
	//输出debug日志
	log.Debug("debug_message", A{a: 1, b: "2"})
	//输出自定义日志
	log.Log(0, logger.DEBUG, []interface{}{"log_message"})

	//设置日志级别为INDO
	log.SetLevel(logger.INFO)
	//设置终端输出的日志界别为INFO
	log.SetConsoleLevel(logger.INFO)
	//打印日志
	log.Debug("debug_message")
	log.Info("info_message")
	log.Warn("warn_message")

	//设置日志回调，可以自行处理日志，例如发送到日志服务器。
	log.SetCallBackFunc(func(f *logger.Format) {
		fmt.Println("Call back", f)
	})

	//打印日志
	log.Error("error_message")
	log.Fatal("fatal_message")

	//未设置默认日志时，只在终端输出日志内容
	logger.Debug("Use default logger debug")
	//未设置默认日志时，输出自定义日志
	logger.Log(0, logger.DEBUG, []interface{}{"log_message"})

	//设置默认日志（在一处初始化并设置为默认，在其他文件内可直接调用logger.xxx打印日志，注意log和logger的区别）
	logger.SetDefaultLogger(log)
	//打印日志
	logger.Debug("debug_message", 2)

	//设置默认日志关闭终端显示
	logger.GetDefaultLogger().SetConsoleLevel(logger.OFF)
	//设置默认日志关闭写入文件
	logger.GetDefaultLogger().SetLevel(logger.OFF)

	//打印日志
	logger.Error("detault_debug_message")
}
