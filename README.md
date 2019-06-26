## logger
Minimalistic logging library for Go.

## Install

```bash
$ go get github.com/goroom/logger
```

## Getting Started
Use default logger
```go
package main

import (
	"github.com/goroom/logger"
)

func main() {
	logger.Debug("No save files, just show at console.")
}
```

Use customer logger
```go
package main

import (
	"fmt"
	"context"
	"github.com/goroom/logger"
)

func main() {
    logger.InitDefaultLogger(
    	logger.WithFileLevel(logger.INFO),
    	logger.WithFileSize(logger.MB/10),
    	logger.WithFileSplit(logger.DefaultFileNoSplit),
    	logger.WithCallBack(func(f *logger.Format) {
    		fmt.Println(f.Time)
    	}),
    )
    logger.Debugf("%d %s", 1, "name")
}
```

## example
```bash
$ cd example
$ go run ./main.go
```