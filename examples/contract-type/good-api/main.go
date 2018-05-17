package main

import "fmt"

func main() {
	l := Logger{"Sample message"}

	// Logging with ConsoleWriter
	cw := ConsoleWriter{}
	l.Log(cw)

	// Logging with TextWriter
	tw := TextWriter{}
	l.Log(tw)

	// Logging with XmlWriter
	xw := XmlWriter{}
	l.Log(xw)
}

type XmlWriter struct{}

// XmlWriter writes logs into xml file
func (x XmlWriter) Write(message string) {
	fmt.Println("Write into xml file...")
	fmt.Println("Log:", message)
}
