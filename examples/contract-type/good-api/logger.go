package main

import "fmt"

type Logger interface {
	Write()
}

type ConsoleLogger struct {
	Message string
}

func (c ConsoleLogger) Write() {
	fmt.Println("Write into console...")
	fmt.Println("Log:", c.Message)
}

type TextLogger struct {
	Message string
}

func (t TextLogger) Write() {
	fmt.Println("Write into text file...")
	fmt.Println("Log:", t.Message)
}

type LoggerHelper struct {
}

func (lh LoggerHelper) Log(l Logger) {
	l.Write()
}
