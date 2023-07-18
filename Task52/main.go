package main

import "fmt"

func isSimetrya(n int, arr *[]int) bool {
	var dl int
	flag := true
	if n%2 == 0 {
		dl = n / 2
	} else {
		dl = n/2 - 1
	}
	for i := 0; i < dl; i++ {
		if (*arr)[i] != (*arr)[n-i-1] {
			flag = !flag
			break
		}
	}
	return flag
}
func main() {
	var n int

	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	if isSimetrya(n, &arr) {
		fmt.Println(0)
	} else {
		if arr[n-2] == arr[n-1] {

		}
	}
}
