package main

import "fmt"

type TextWriter struct{}

// TextWriter writes logs into text file
func (t TextWriter) Write(message string) {
	fmt.Println("Write into text file...")
	fmt.Println("Log:", message)
}

type Logger struct {
	Message string
}

func (l Logger) Log(t TextWriter) {
	t.Write(l.Message)
}
