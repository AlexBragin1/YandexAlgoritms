package main

import "fmt"

func main() {
	var n, k, num int

	fmt.Scanln(&n, &k)
	arrN := make([]int, 0, n+1)
	arrN = append(arrN, 1)
	for i := 1; i <= n; i++ {

		if i < k {
			for j := 0; j < len(arrN); j++ {
				num += arrN[j]

			}
		} else {
			for j := i - k; j < len(arrN); j++ {
				num += arrN[j]

			}
		}
		arrN = append(arrN, num)
		num = 0
	}

	fmt.Println(arrN[n-1])
}
