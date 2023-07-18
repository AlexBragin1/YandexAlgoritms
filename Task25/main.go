package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, a int
	fmt.Scanln(&n)
	arr := make([]int, 0, n)
	dp := make([]int, 0, n+1)
	for i := 0; i < n; i++ {

		fmt.Scan(&a)
		arr = append(arr, a)
	}
	sort.Ints(arr)

	dp = append(dp, arr[1]-arr[0])
	if n > 2 {
		dp = append(dp, arr[2]-arr[0])

		for i := 3; i < n; i++ {
			if dp[i-3] < dp[i-2] {
				dp = append(dp, dp[i-3]+arr[i]-arr[i-1])
			} else {
				dp = append(dp, dp[i-2]+arr[i]-arr[i-1])
			}
		}
	}
	fmt.Println(dp[n-2])
}
