package main

func main() {
	l := TextLogger{"Sample message"}
	lh := LoggerHelper{}
	lh.Log(l)
}
