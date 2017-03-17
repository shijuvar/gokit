package main

import (
	"errors"
	"flag"

	"github.com/shijuvar/gokit/log"
)

func main() {
	// Parse log level from command line
	logLevel := flag.Int("loglevel", 0, "an integer value (0-4)")
	flag.Parse()
	// Calling the SetLogLevel with the command-line argument
	log.SetLogLevel(log.Level(*logLevel), "logs.txt")
	log.Trace.Println("Main started")
	loop()
	err := errors.New("Sample Error")
	log.Error.Println(err.Error())
	log.Trace.Println("Main completed")
}

// A simple function for the logging demo
func loop() {
	log.Trace.Println("Loop started")
	for i := 0; i < 10; i++ {
		log.Info.Println("Counter value is:", i)
	}
	log.Warning.Println("The counter variable is not being used")
	log.Trace.Println("Loop completed")
}
