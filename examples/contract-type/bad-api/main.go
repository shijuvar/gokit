package main

func main() {
	tw := TextWriter{"Sample message"}
	l := Logger{}
	// Log with TextWriter
	l.Log(tw)
}
