package lib

import (
	"fmt"
)

// Print all favorites
func PrintFavorites() {
	for _, v := range favorites {
		fmt.Println(v)
	}
}
