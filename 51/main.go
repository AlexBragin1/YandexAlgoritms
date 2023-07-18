package main

import "fmt"

func main() {
	var n, a int
	fmt.Scanln(&n)
	m := make([]int, 0, n)
	min := 0

	for i := 0; i < n; i++ {

		fmt.Scanln(&a)
		m = append(m, a)
	}

	for i := 0; i < n-1; i++ {
		if m[i] <= m[i+1] {
			min += m[i]
		} else {
			min = m[i+1]
		}
		fmt.Println(min)
	}
	fmt.Println(min)
}
