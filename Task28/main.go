package main

import "fmt"

func main() {
	var m, n int
	fmt.Scanln(&m, &n)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			dp[i] = append(dp[i], 0)
		}
	}

	dp[1][1] = 1

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-2][j-1] + dp[i-1][j-2]
		}
	}
	fmt.Println(dp[m][n])

}
