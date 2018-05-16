package main

import "fmt"

type Writer interface {
	Write()
}

type ConsoleWriter struct {
	Message string
}

// ConsoleWriter writes logs into stdout
func (c ConsoleWriter) Write() {
	fmt.Println("Write into console...")
	fmt.Println("Log:", c.Message)
}

type TextWriter struct {
	Message string
}

// TextWriter writes logs into text file
func (t TextWriter) Write() {
	fmt.Println("Write into text file...")
	fmt.Println("Log:", t.Message)
}

type Logger struct {
}

func (lh Logger) Log(writer Writer) {
	writer.Write()
}
