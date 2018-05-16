package main

import "fmt"

type TextWriter struct {
	Message string
}

func (t TextWriter) Write() {
	fmt.Println("Write into text file...")
	fmt.Println("Log:", t.Message)
}

type Logger struct {
}

func (l Logger) Log(t TextWriter) {
	t.Write()
}
