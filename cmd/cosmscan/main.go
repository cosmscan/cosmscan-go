package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("Cosmscan started")
}
