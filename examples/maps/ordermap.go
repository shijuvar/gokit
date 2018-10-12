package main

import "fmt"

// orderMap is a data structure that preserves the order of the data insert into map.
type orderMap struct {
	store map[string]string
	order []string
}

func newOrderMap() *orderMap {
	return &orderMap{
		store: make(map[string]string),
		order: make([]string, 0),
	}
}

func (om *orderMap) insert(k, v string) {
	om.store[k] = v
	om.order = append(om.order, k)
}

func (om *orderMap) printAll() {
	for _, v := range om.order {
		fmt.Println(om.store[v])
	}
}

func main() {
	ms := newOrderMap()
	ms.insert("jpmb", "JP Morgan, Bangalore")
	ms.insert("jpmh", "JP Morgan, Hyderabad")
	ms.printAll()

}
