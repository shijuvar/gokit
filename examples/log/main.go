package main

import (
	"errors"
	"flag"
)

func main() {
	// Parse log level from command line
	logLevel := flag.Int("loglevel", 0, "an integer value (0-4)")
	flag.Parse()
	// Calling the SetLogLevel with the command-line argument
	SetLogLevel(Level(*logLevel))
	Trace.Println("Main started")
	loop()
	err := errors.New("Sample Error")
	Error.Println(err.Error())
	Trace.Println("Main completed")
}

// A simple function for the logging demo
func loop() {
	Trace.Println("Loop started")
	for i := 0; i < 10; i++ {
		Info.Println("Counter value is:", i)
	}
	Warning.Println("The counter variable is not being used")
	Trace.Println("Loop completed")
}
