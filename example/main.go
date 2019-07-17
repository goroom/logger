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
		logger.WithFileLevel(logger.ALL),
		logger.WithFileSplit(logger.DefaultFileSplitByMinute),
		logger.WithCallBack(func(f *logger.Format) {
			//fmt.Println(f.Time)
		}),
	)

	for i := 0; i < 100; i++ {
		logger.Debug(time.Now().UnixNano())
		logger.Info(time.Now().UnixNano())
		logger.Warn(time.Now().UnixNano())
		logger.Error(time.Now().UnixNano())
		//
		logger.Debugf("Debugf %v", i)
		logger.Infof("Infof %v", i)
		logger.Warnf("Warnf %v", i)
		logger.Errorf("Errorf %v", i)
		time.Sleep(time.Second * 1)
	}
	logger.Wait()
}
