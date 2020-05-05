package main

import (
	"fmt"
)

type Writer interface {
	Write(string)
}

// func type
// Refer http package's HandlerFunc func type
// http.HandlerFunc is an adapter to func(http.ResponseWriter, *http.Request)
// http.HandlerFunc is also an implementation of http.Handler interface
// http.Handler interface defines one behaviour: ServeHTTP(http.ResponseWriter, *http.Request)
type StringWriter func(string) string

// Implements Writer interface
func (s StringWriter) Write(str string) {
	fmt.Println(s(str))
}

// XStringer can plug into StringWriter adapter
func XStringer(str string) string {
	return "XStringer: " + str
}

// YStringer can plug into StringWriter adapter
func YStringer(str string) string {
	return "YStringer: " + str
}

// Prints variadic values of StringWriter func type
func print(str string, fs ...StringWriter) {
	for _, v := range fs {
		v.Write(str)
	}
}

// Prints variadic values of Writer interface
func printInterfaceValues(str string, fs ...Writer) {
	for _, v := range fs {
		v.Write(str)
	}
}
func main() {

	z := func(str string) string {
		return "ZStringer: " + str
	}

	// Map of func values of interface type
	fni := map[string]Writer{
		"X": StringWriter(XStringer),
		"Y": StringWriter(YStringer),
		"Z": StringWriter(z),
	}

	// Call Write method
	for _, v := range fni {
		v.Write("test with interface value")
	}
	fmt.Println("*************************")
	// Map of func values
	fnf := map[string]StringWriter{
		"X": StringWriter(XStringer),
		"Y": StringWriter(YStringer),
		"Z": StringWriter(z),
	}
	// Call Write method
	for _, v := range fnf {
		v.Write("test with func value")
	}
	fmt.Println("*************************")

	fs := []StringWriter{
		StringWriter(XStringer),
		StringWriter(YStringer),
	}

	fs = append(fs,
		StringWriter(func(str string) string { //In-line func value
			return "In-line Stringer: " + str
		}),
	)

	// func value to a variable
	zStringer := func(str string) string {
		return "ZStringer: " + str
	}
	fs = append(fs, zStringer)
	//print
	print("test on variadic func values", fs...)
}
