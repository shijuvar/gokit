package main

import (
	"context"
	"log/slog"
	"os"
)

type Order struct {
	ID         int
	CustomerID string
}

func main() {
	slog.Info("hello, world")
	slog.Info("hello, world", "user", os.Getenv("USER"))
	// slog’s top-level functions use the default logger.
	// We can get this logger explicitly,
	logger := slog.Default()
	order := Order{
		ID:         100,
		CustomerID: "shijuvar",
	}
	doLog(logger, order)
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	doLog(logger, order)
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	doLog(logger, order)

}

func doLog(logger *slog.Logger, o Order) {
	logger.Info("New Order has been created", "Order ID", o.ID, "CustomerID", o.CustomerID)
	// With Attr type. Attr is a key-value pair.
	logger.Info("Order", slog.Int("Order ID", o.ID), slog.String("CustomerID", o.CustomerID))

	logger.LogAttrs(context.Background(), slog.LevelInfo, "Order is created by",
		slog.String("user", os.Getenv("USER")))
}
