package main

import "fmt"

func main() {
	var n int
	var a, b, c int
	var min int
	fmt.Scan(&n)
	dp := make([]int, n+3)
	matrix := make([][]int, n+3)
	for i := 0; i < n+3; i++ {

		if i < 3 {
			matrix[i] = append(matrix[i], 0)
			matrix[i] = append(matrix[i], 0)
			matrix[i] = append(matrix[i], 0)
		} else {
			fmt.Scanln(&a, &b, &c)
			matrix[i] = append(matrix[i], a)
			matrix[i] = append(matrix[i], b)
			matrix[i] = append(matrix[i], c)

		}

	}

	dp[0], dp[1], dp[2] = 0, 0, 0
	for i := 3; i < n+3; i++ {
		min = dp[i-1] + matrix[i][0]

		if min > dp[i-2]+matrix[i-1][1] && dp[i-2]+matrix[i-1][1] > 0 {
			min = dp[i-2] + matrix[i-1][1]
		}
		if min > dp[i-3]+matrix[i-2][2] && matrix[i-2][2] > 0 {
			min = dp[i-3] + matrix[i-2][2]
		}
		dp[i] = min

	}
	fmt.Println(dp[n+2])
}
