package zaplogger

import (
	"time"

	zap "go.uber.org/zap"
)

func Run() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Info("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "url",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "url")
}
