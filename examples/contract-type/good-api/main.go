package main

import "fmt"

func main() {
	lh := LoggerHelper{}

	// Logging with ConsoleLogger
	cl := ConsoleLogger{"Sample message"}
	lh.Log(cl)

	// Logging with TextLogger
	tl := TextLogger{"Sample message"}
	lh.Log(tl)

	// Logging with XmlLogger
	xl := XmlLogger{"Sample message"}
	lh.Log(xl)
}

type XmlLogger struct {
	Message string
}

func (x XmlLogger) Write() {
	fmt.Println("Write into xml file...")
	fmt.Println("Log:", x.Message)
}
