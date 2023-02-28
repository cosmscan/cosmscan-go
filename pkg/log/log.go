package log

import (
	"errors"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once

// Logger is the global logger used in the application
var Logger = zap.S()

var (
	ErrInvalidEncoding     = errors.New("invalid encoding, supported encodings are 'console' and 'json'")
	ErrUnsupportedLogLevel = errors.New("unsupported log level")
)

// InitLogger sets the global logger
func InitLogger(cfg Config) error {
	var err error
	var logger *zap.Logger

	once.Do(func() {
		if cfg.Encoding != "console" && cfg.Encoding != "json" {
			err = ErrInvalidEncoding
			return
		}

		lv, pErr := zapcore.ParseLevel(cfg.Level)
		if pErr != nil {
			err = errors.Join(pErr, ErrUnsupportedLogLevel)
			return
		}

		c := zap.NewProductionConfig()
		c.Encoding = cfg.Encoding
		c.Level = zap.NewAtomicLevelAt(lv)

		logger, err = c.Build()
		if err != nil {
			return
		}

		defer logger.Sync()
		zap.ReplaceGlobals(logger)
		Logger = logger.Sugar()
	})

	return err
}
