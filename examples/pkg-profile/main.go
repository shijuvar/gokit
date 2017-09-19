package main

import (
	"github.com/pkg/profile"
	"math"
)

func main() {

	defer profile.Start(profile.CPUProfile).Stop()
	pow()
}

func pow() {
	for i := 1; i < 10000; i++ {
		math.Pow10(i)
	}
}
