package main

import (
	"errors"
	"fmt"
)

var data map[string]string

// init function will be automatically invoked before the main function
func init() {
	data = make(map[string]string) // Initialise map with make
}

func getByKey(key string) (string, error) {
	if val, ok := data[key]; ok {
		return val, nil
	} else {
		//return "", errors.New("key does not exist")
		return "", fmt.Errorf("key: %s,  not found", key)
	}
}
func addItem(k, v string) error {
	if _, ok := data[k]; ok {
		//return fmt.Errorf("%s already exists", k)
		return errors.New("key already exists")
	}
	data[k] = v // insert into map
	return nil
}

func main() {
	if err := addItem("HI", "Hindi"); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("data has been added")

	if v, err := getByKey("HId"); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(v)
	}

}
