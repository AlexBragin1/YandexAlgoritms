package main

import (
	"fmt"
)

func main() {
	var n, k1, k2, min int
	var kuponDay []int
	fmt.Scanln(&n)

	if n == 0 {
		fmt.Println(0)
		fmt.Println(0, 0)
	} else {
		priceObed := make([]int, n)
		dp := make([][]int, n)
		fmt.Scanln(&priceObed[0])
		for i := 0; i < n; i++ {

			dp[i] = append(dp[i], 300001)

			if i == 0 && priceObed[0] <= 100 {
				dp[i] = append(dp[i], priceObed[0])
			}
			if i == 0 && priceObed[0] > 100 {
				dp[i] = append(dp[i], 300001)
				dp[i] = append(dp[i], priceObed[0])
				continue
			}
			if i < (n - 1) {
				dp[0] = append(dp[0], 300001)
			}

		}

		for i := 1; i < n; i++ {
			fmt.Scanln(&priceObed[i])
			for j := 1; j <= n; j++ {
				if priceObed[i] > 100 {
					if dp[i-1][j-1] != 300001 {
						min = dp[i-1][j-1] + priceObed[i]
					} else {
						min = 300001
					}
					if j != n && dp[i-1][j+1] < 300001 && min > dp[i-1][j+1] {
						min = dp[i-1][j+1]
					}

				} else {

					if dp[i-1][j] != 300001 {
						min = dp[i-1][j] + priceObed[i]
					} else {
						min = 300001
					}
					if j != n && dp[i-1][j+1] < 300001 && min > dp[i-1][j+1] {
						min = dp[i-1][j+1]
					}
				}
				dp[i] = append(dp[i], min)
			}

		}
		min = dp[n-1][1]

		for i := 1; i < len(dp[n-1]); i++ {
			if min >= dp[n-1][i] {
				min = dp[n-1][i]
				k1 = i - 1

			}

		}

		j := k1 + 1

		for i := n - 1; i > 0; i-- {
			if dp[i][j] == dp[i-1][j+1] {
				if j+1 > k2 {
					k2++
					j = j + 1

					kuponDay = append(kuponDay, i+1)

				}
			} else {
				if priceObed[i] > 100 {
					j = j - 1
				}

			}
		}

		fmt.Println(min)
		fmt.Println(k1, k2)

		for i := len(kuponDay) - 1; i >= 0; i-- {
			fmt.Println(kuponDay[i], " ")
		}
		for i, _ := range dp {
			fmt.Println(dp[i])
		}
	}
}
