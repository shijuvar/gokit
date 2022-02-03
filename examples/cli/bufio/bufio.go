package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter value:")
	val, _ := reader.ReadString('\n')
	fmt.Println("The value is", val)
}
