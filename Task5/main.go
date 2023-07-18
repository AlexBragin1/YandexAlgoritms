package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scanln(&n)

	min := 0
	fmt.Scanln(&a)
	for i := 1; i < n; i++ {

		fmt.Scanln(&b)
		if a < b {
			min += a
		} else {
			min += b
		}
		a = b
	}

	fmt.Println(min)
}
