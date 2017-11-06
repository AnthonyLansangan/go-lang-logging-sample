package main

import (
	"errors"
	"lumberjack_sample/logger"
)

func main() {

	logger.Info("This is a sample log")
	logger.Warning("Warning log")
	logger.Error(errors.New("Encountered an error"))
	logger.Trace("Showing traces")
	logger.Fatal("System will exit")
}
