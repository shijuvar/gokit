package main

import (
	"errors"
	"fmt"

	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

func main() {

	block := make(chan struct{})
	result1 := getTemperature("kochi", block)
	result2 := getTemperature("kochi", block)
	result3 := group.DoChan("kochi", func() (interface{}, error) {
		<-block
		return 25, nil
	})
	result4 := getTemperature("bengaluru", block)

	close(block)
	r1 := <-result1
	r2 := <-result2
	r3 := <-result3
	r4 := <-result4
	fmt.Println("r1:", r1.Val)
	fmt.Println("r2:", r2.Val)
	fmt.Println("r3:", r3.Val)
	fmt.Println("r4:", r4.Val)

	// Results are shared by functions executed with duplicate keys.
	fmt.Printf("Shared r1:%v, r2:%v, r3:%v, r4:%v", r1.Shared, r2.Shared, r3.Shared, r4.Shared)
}

func getTemperature(city string, block chan struct{}) <-chan singleflight.Result {
	result := group.DoChan(city, func() (interface{}, error) {
		<-block
		tempC, err := getWeather(city)
		return tempC, err
	})
	return result
}

func getWeather(city string) (int, error) {
	// hardcoded results just for the sake of example
	switch city {
	case "kochi":
		return 28, nil
	case "bengaluru":
		return 22, nil
	default:
		return 0, errors.New("city not found")
	}
	return 0, errors.New("error")
}
