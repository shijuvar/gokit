package main

import "fmt"

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

func (l Logger) Log(t TextWriter) {
	t.Write()
}
