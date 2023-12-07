package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type Order struct {
	ID         int
	CustomerID string
}

// withOldLogAPIs demonstrates how to use slog with older API which need log.Logger
func withOldLogAPIs() {
	// NewJSONHandler creates a JSONHandler
	handler := slog.NewJSONHandler(os.Stdout, nil)
	// NewLogLogger returns a new log.Logger
	// The logger acts as a bridge from the older log API to newer structured logging handlers
	logger := slog.NewLogLogger(handler, slog.LevelError)
	server := http.Server{
		Addr:     ":8080",
		Handler:  nil,
		ErrorLog: logger,
	}
	server.ListenAndServe()
	defer server.Close()
}
func logWithLevels() {
	// log using the default logger
	slog.Debug("Debug log")
	slog.Info("Info log")
	slog.Warn("Warning log")
	slog.Error("Error log")
}
func main() {
	logWithLevels()
	// Write log as key-value pairs
	// accept a log message as their first argument,
	// and variadic number of key/value pairs
	slog.Info("Order has been created", "OrderID", "10001")
	// !BADKEY
	slog.Info("Order has been created", "10001")
	// slog’s top-level functions use the default logger.
	// We can get this logger explicitly,
	logger := slog.Default() // Default returns the default Logger.
	order := Order{
		ID:         100,
		CustomerID: "shijuvar",
	}
	doLog(logger, order)
	// New creates a new Logger with the given non-nil Handler.
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	doLog(logger, order)
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	doLog(logger, order)
	customizeHandler()

}

func doLog(logger *slog.Logger, o Order) {
	logger.Info("New Order has been created", "Order ID", o.ID, "CustomerID", o.CustomerID)
	// With slog.Attr type. slog.Attr is a key-value pair.
	logger.Info("Order", slog.Int("Order ID", o.ID), slog.String("CustomerID", o.CustomerID))
	// LogAttrs is a more efficient version of Logger.Log that accepts only Attrs.
	logger.LogAttrs(context.Background(), slog.LevelInfo, "Order has been created",
		slog.String("created by", os.Getenv("USER")))

	// Grouping contextual attributes
	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"Order has been created",
		slog.String("created by", os.Getenv("USER")),
		slog.Group("Order",
			slog.Int("Order ID", o.ID),
			slog.String("CustomerID", o.CustomerID),
		),
	)
}

func customizeHandler() {
	fmt.Println("-------customizeHandler-------")
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug, // Levels: DEBUG (-4), INFO (0), WARN (4), and ERROR (8).
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	logger.Debug("debugging message")
}
