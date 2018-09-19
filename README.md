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
    logger.SetFileLevel(logger.ALL)       // default off
    logger.SetFileSize(logger.KB.CalB(2)) // default 5MB
    logger.SetFileCount(3)                // default 5
    logger.SetFileBaseName("ex")          // default exec name
    logger.SetContextCallBackFunc(func(ct context.Context) string {
    	return "id"
    })
    logger.SetConsoleFormat(func(f *logger.Format) string {
    	return fmt.Sprintf("%s %v", f.Time.Format("15:02:03"), f.ArgsDefaultFormat())
    })
    logger.SetCallBackFunc(func(f *logger.Format) {
		fmt.Println(f)
	})
    logger.DebugF("%d %s", 1, "name")
}
```

## example
```bash
$ cd example
$ go run ./main.go
```