package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"time"
)

func main() {
	typeAndValue()
	formatToString()
}

func typeAndValue() {
	num := 100
	str := "Gopher"
	// TypeOf returns the reflection Type that represents the dynamic type.
	tNum := reflect.TypeOf(num)
	tStr := reflect.TypeOf(str)
	fmt.Println("-----reflect.Type-----")
	fmt.Println(tNum)
	fmt.Println(tStr)
	// reflect.TypeOf returns interface value's concrete type
	var r io.Reader = os.Stdin
	var f interface{} = float64(100)
	fmt.Println(reflect.TypeOf(r))
	fmt.Println(reflect.TypeOf(f))
	// ValueOf returns a new Value initialized to the concrete value
	v := reflect.ValueOf(num)
	fmt.Println("-----reflect.Value-----")
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
	fmt.Println(v.Kind()) // Kind returns v's Kind.

	// Method Type returns reflect.Type as its type
	t := v.Type()
	fmt.Println("-----Value.Type-----")
	fmt.Println(t.String())
	v = reflect.ValueOf(num)
	// a reflect.Value
	x := v.Interface()
	// Returns an interface{}
	i := x.(int)
	// type assertion to int
	fmt.Println("-----Value.Interface-----")
	fmt.Printf("%d\n", i)
}

func formatToString() {
	var x int64 = 100
	var d time.Duration = 10 * time.Nanosecond
	fmt.Println("----Formats any value as a string----")
	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any([]int64{x}))
	fmt.Println(Any([]time.Duration{d}))

}

// Any formats any value as a string.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
		// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
