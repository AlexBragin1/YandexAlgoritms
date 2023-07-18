package main

import "fmt"

func main() {
	var m, n, result int
	var miniA, maxiB, minjA, maxjB int
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	a, b := make([]int, n), make([]int, n)
	c := make([]bool, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&a[i], &b[i])
		c[i] = true
		for j := 0; j < i; j++ {

			if a[i] < b[i] {
				miniA, maxiB = a[i], b[i]
			} else {
				miniA, maxiB = b[i], a[i]
			}
			if a[j] < b[j] {
				minjA, maxjB = a[j], b[j]
			} else {
				minjA, maxjB = b[j], a[j]
			}

			if miniA <= maxjB && minjA <= maxiB {
				c[j] = false
			}

		}

	}

	for i := 0; i < n; i++ {
		if c[i] {
			result++
		}
	}

	fmt.Println(result)

}
