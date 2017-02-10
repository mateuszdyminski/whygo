package main

import (
	"github.com/uber-go/zap"
	"os"
	"path"
)

func main() {
	logger := zap.New(
		zap.NewJSONEncoder(zap.RFC3339Formatter("key")),
		zap.Fields(zap.Int("pid", os.Getpid()),
			zap.String("exe", path.Base(os.Args[0]))),
	)

	logger.Info("Hello Wroclaw")
}
