package main

import "fmt"

type Writer interface {
	Write(string)
}

type ConsoleWriter struct{}

// ConsoleWriter writes logs into stdout
func (c ConsoleWriter) Write(message string) {
	fmt.Println("Write into console...")
	fmt.Println("Log:", message)
}

type TextWriter struct{}

// TextWriter writes logs into text file
func (t TextWriter) Write(message string) {
	fmt.Println("Write into text file...")
	fmt.Println("Log:", message)
}

type Logger struct {
	Message string
}

func (l Logger) Log(writer Writer) {
	writer.Write(l.Message)
}
