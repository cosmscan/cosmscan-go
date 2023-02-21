package log

import (
	"sync"

	"go.uber.org/zap"
)

var once sync.Once

// Logger is the global logger used in the application
var Logger = zap.S()

func init() {
	once.Do(func() {
		logger, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		defer logger.Sync()
		zap.ReplaceGlobals(logger)
		Logger = logger.Sugar()
	})
}
