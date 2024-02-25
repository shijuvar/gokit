package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
)

type Order struct {
	ID         int
	CustomerID string
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

/*
If a type implements the LogValuer interface,
the Value returned from its LogValue method is used for logging

	type LogValuer interface {
		LogValue() slog.Value
	}
*/

func (u User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", u.ID),
		slog.String("name", u.FirstName+" "+u.LastName),
	)
}

func logWithLogValuer() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)

	u := User{
		ID:        "shijuvar",
		FirstName: "shiju",
		LastName:  "varghese",
		Password:  "my-password",
	}

	logger.Info("info", "user", u)
}

// logWithLevels demonstrates basic logging all levels
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
	logger.Info("Order has been created", "order", order)
	// logging with different approaches
	fmt.Println("Logging with default logger")
	doLog(logger, order)
	// New creates a new Logger with the given non-nil Handler.
	fmt.Println("Logging with TextHandler")
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	doLog(logger, order)
	fmt.Println("Logging with JSONHandler")
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	doLog(logger, order)
	// customizing log handler by providing slog.HandlerOptions value
	customizeHandler()
	// logging with LogValuer interface to hide sensitive data
	logWithLogValuer()
}

// doLog demonstrates logging with different ways which includes slog.Attr type, LogAttrs, slog.Group
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

// customizeHandler customizes log handler by providing slog.HandlerOptions value
func customizeHandler() {
	fmt.Println("-------customizeHandler-------")
	// AddSource causes the handler to compute the source code position of the log statement
	// and add a SourceKey attribute to the output.
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug, // Levels: DEBUG (-4), INFO (0), WARN (4), and ERROR (8).
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	logger.Debug("debugging message")
}

// withOldLogAPIs demonstrates how to use slog with older API which need log.Logger
func withOldLogAPIs() {
	// NewJSONHandler creates a JSONHandler
	handler := slog.NewJSONHandler(os.Stdout, nil)
	// NewLogLogger returns a new log.Logger
	// The logger acts as a bridge from the older log API to newer structured logging http
	logger := slog.NewLogLogger(handler, slog.LevelError)
	server := http.Server{
		Addr:     ":8080",
		Handler:  nil,
		ErrorLog: logger,
	}
	server.ListenAndServe()
	defer server.Close()
}

// zapAsTheBackEnd demonstrates how to use Uber Zap as the backend for slog
func zapAsTheBackEnd() {
	zapLogger := zap.Must(zap.NewProduction())

	defer zapLogger.Sync()
	// creating a slog frontend
	logger := slog.New(zapslog.NewHandler(zapLogger.Core(), nil))

	logger.Info(
		"New Order has been created",
		slog.Int("Order ID", 10001),
		slog.String("CustomerID", "shijuvar"),
	)
}
