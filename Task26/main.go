package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var m, n int
	value, _ := reader.ReadString('\n')
	value = strings.Trim(value, "\n")
	mas := strings.Split(value, " ")
	m, _ = strconv.Atoi(mas[0])
	n, _ = strconv.Atoi(mas[1])
	matrix := make([][]int, m)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		value, _ = reader.ReadString('\n')
		value = strings.Trim(value, "\n")
		mas = strings.Split(value, " ")

		for j := 0; j < n; j++ {
			temp, _ := strconv.Atoi(mas[j])
			matrix[i] = append(matrix[i], temp)
			dp[i] = append(dp[i], 0)
		}
	}
	dp[0][0] = matrix[0][0]
	for j := 1; j < n; j++ {
		dp[0][j] = matrix[0][j] + dp[0][j-1]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = matrix[i][0] + dp[i-1][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + matrix[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + matrix[i][j]
			}
		}
	}

	fmt.Println(dp[m-1][n-1])

}
