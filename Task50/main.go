package main

import (
	"fmt"
)

func main() {
	var a, b, c, x int
	var result string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	if a == 0 {
		result = "0"
	}
	if c < 0 {
		result = "NO SOLUTION "
	}

	fmt.Println(result)
}
