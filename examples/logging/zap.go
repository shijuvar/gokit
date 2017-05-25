package main

import (
	"time"

	"go.uber.org/zap"
)

const url = "medium.com/@shijuvar"

func main() {
	SugaredLogger()
	Logger()
}
func SugaredLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	// flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL.",
		// Structured context as loosely-typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	// Infof uses fmt.Sprintf to log a templated message
	sugar.Infof("Failed to fetch URL: %s", url)
}
func Logger() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("Failed to fetch URL.",
		// Structured context as strongly-typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
