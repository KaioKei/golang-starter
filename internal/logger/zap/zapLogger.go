package zap

import (
	"go.uber.org/zap"
	"time"
)

func Start() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushed buffer after function return, if any
	url := "http://issou.com"

	// Standard logger
	// Log headers will be injected like the log level, timestamp, caller (file origin), then msg
	// For more than a message content and add key:value information, you can use 'zap.<type>' entries
	// inside the log level function's paramaters
	logger.Info("Hello world !")
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	// Sugar logger is a logger type that allows nice coding style but is slightly slower than
	// the standard one.
	// Using the sugar Infow, you can refer to a key:value format
	// Using the sugar Infof, you can refer to a string formatted style like a Printf style
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
