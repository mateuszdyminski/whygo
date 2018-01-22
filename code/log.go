package main

import (
	"os"
	"path"

	"go.uber.org/zap"
)

func main() {
	logger := zap.New(
		zap.NewJSONEncoder(zap.RFC3339Formatter("key")),
		zap.Fields(zap.Int("pid", os.Getpid()),
			zap.String("exe", path.Base(os.Args[0]))),
	)

	logger.Info("Hello Wroclaw")
}
