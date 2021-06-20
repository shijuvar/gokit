// Run here: https://go2goplay.golang.org
package main

import (
	"fmt"
)

//func sliceOf(type T)(elems ...T) []T {
//	genSlice := []T{}
//	for _, el := range elems {
//		arr = append(arr, el)
//	}
//
// return genSlice
//}
type Stringer interface {
	ToString() string
}

type person struct {
	Fname, Lname string
}

func (p person) ToString() string {
	return fmt.Sprint(p.Fname, " ", p.Lname)
}

//func stringify(type T Stringer)(s []T) (ret []string) {
//for _, v := range s {
//ret = append(ret, v.ToString())
//}
//return ret
//}

func main() {
	//strs := sliceOf("one", "two", "three", "four", "five")
	//fmt.Println(strs)
	//
	//ints := sliceOf(1, 2, 3, 4, 5)
	//fmt.Println(ints)
	//p := []Stringer{
	//	person{"shiju", "varghese"},
	//	person{"irene", "rose"},
	//}
	//fmt.Println(p)
	//s := stringify(p)
	//fmt.Println(s)
}
