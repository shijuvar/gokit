package main

import (
	"fmt"
)

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

func (l Level) String() string {
	return [...]string{"UNSPECIFIED", "TRACE", "INFO", "WARNING", "ERROR"}[l]
}
func main() {
	var l Level = TRACE
	fmt.Println(l.String())

}
