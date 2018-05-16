package main

func main() {
	tw := TextWriter{"Sample message"}
	l := Logger{}
	l.Log(tw)
}
