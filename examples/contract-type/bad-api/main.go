package main

func main() {
	tw := TextWriter{}
	l := Logger{"Sample message"}
	// Log with TextWriter
	l.Log(tw)
}
