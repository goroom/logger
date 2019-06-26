package main

import (
	"os"
	"path"
	"time"

	"github.com/goroom/logger"
)

func main() {
	logger.GetDefaultLogger().Info(path.Base(os.Args[0]))

	logger.InitDefaultLogger(
		logger.WithFileLevel(logger.INFO),
		logger.WithFileSize(logger.MB/10),
		logger.WithFileSplit(logger.DefaultFileNoSplit),
		logger.WithCallBack(func(f *logger.Format) {
			//fmt.Println(f.Time)
		}),
	)

	for i := 0; i < 10000; i++ {
		logger.Debug(time.Now().UnixNano())
		logger.Info(time.Now().UnixNano())
		logger.Warn(time.Now().UnixNano())
		logger.Error(time.Now().UnixNano())
		//
		logger.Debugf("Debugf %v", i)
		logger.Infof("Infof %v", i)
		logger.Warnf("Warnf %v", i)
		logger.Errorf("Errorf %v", i)
		time.Sleep(time.Millisecond * 1)
	}
	logger.Wait()
}
