package main

import (
	"flag"
	"fmt"
)

func main() {

	fileName := flag.String("filename", "logfile", "File name for the log file")
	logLevel := flag.Int("loglevel", 0, "An integer value for Level (0-4)")
	isEnable := flag.Bool("enable", false, "A boolean value for enabling log options")
	var num int
	// Bind the flag to a variable.
	flag.IntVar(&num, "num", 25, "An integer value")

	// Parse parses flag definitions from the argument list.
	flag.Parse()
	// Get the values from pointers
	fmt.Println("filename:", *fileName)
	fmt.Println("loglevel:", *logLevel)
	fmt.Println("enable:", *isEnable)
	// Get the value from a variable
	fmt.Println("num:", num)
	// Args returns the non-flag command-line arguments.
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println("The non-flag command-line arguments are:")
		// Print the arguments
		for _, v := range args {
			fmt.Println(v)
		}
	}

}
