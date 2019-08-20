package main

import "fmt"

type Writer interface {
	Write(string)
}

type StringWriter func(string) string

func (s StringWriter) Write(str string) {
	fmt.Println(s(str))
}

func XStringer(str string) string {
	return "XStringer: " + str
}

func YStringer(str string) string {
	return "YStringer: " + str
}

func Print(fs ...StringWriter) {
	for _, v := range fs {
		v.Write("test from variadic func values")
	}
}
func main() {
	// Map of func values of interface type
	fni := map[string]Writer{
		"X": StringWriter(XStringer),
		"Y": StringWriter(YStringer),
	}

	// Call Write method
	for _, v := range fni {
		v.Write("test with interface value")
	}
	// Map of func values
	fnf := map[string]StringWriter{
		"X": StringWriter(XStringer),
		"Y": StringWriter(YStringer),
	}
	// Call Write method
	for _, v := range fnf {
		v.Write("test with func value")
	}

	fs := []StringWriter{
		StringWriter(XStringer),
		StringWriter(YStringer),
	}
	//Print
	Print(fs...)
}
