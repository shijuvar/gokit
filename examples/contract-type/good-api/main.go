package main

import "fmt"

func main() {
	l := Logger{}

	// Logging with ConsoleWriter
	cw := ConsoleWriter{"Sample message"}
	l.Log(cw)

	// Logging with TextWriter
	tw := TextWriter{"Sample message"}
	l.Log(tw)

	// Logging with XmlWriter
	xw := XmlWriter{"Sample message"}
	l.Log(xw)
}

type XmlWriter struct {
	Message string
}

func (x XmlWriter) Write() {
	fmt.Println("Write into xml file...")
	fmt.Println("Log:", x.Message)
}
