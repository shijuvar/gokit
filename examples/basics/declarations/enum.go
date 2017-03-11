package main

import "fmt"

const (
	// UNSPECIFIED logs nothing
	UNSPECIFIED Level = iota // 0 :
	// TRACE logs everything
	TRACE // 1
	// INFO logs Info, Warnings and Errors
	INFO // 2
	// WARNING logs Warning and Errors
	WARNING // 3
	// ERROR just logs Errors
	ERROR // 4
)

// Level holds the log level.
type Level int

// levels provides the string name of Level
var levels = [...]string{
	"UNSPECIFIED",
	"TRACE",
	"INFO",
	"WARNING",
	"ERROR",
}

// String returns the string value of level
func (l Level) String() string {
	return levels[l]
}

func main() {
	level := TRACE
	if level == TRACE {
		fmt.Println("TRACE")
	}
	level = INFO
	fmt.Println(level.String())
}
