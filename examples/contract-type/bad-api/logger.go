package main

import "fmt"

type TextLogger struct {
	Message string
}

func (t TextLogger) Write() {
	fmt.Println("Write into text file...")
	fmt.Println("Log:", t.Message)
}

type LoggerHelper struct {
}

func (lh LoggerHelper) Log(t TextLogger) {
	t.Write()
}
