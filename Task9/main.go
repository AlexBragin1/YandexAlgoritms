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
	var masS []int

	value, _ := reader.ReadString('\n')
	value = strings.Trim(value, "\n")
	mas := strings.Split(value, " ")

	m, _ := strconv.Atoi(mas[0])
	n, _ := strconv.Atoi(mas[1])
	z, _ := strconv.Atoi(mas[2])
	matrix := make([][]int, m)
	sum := make([][]int, m+1)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			sum[i] = append(sum[i], 0)
		}
	}
	for i := 0; i < m; i++ {
		value, _ = reader.ReadString('\n')
		value = strings.Replace(value, "\n", "", -1)
		mas = strings.Split(value, " ")
		for j := 0; j < n; j++ {

			matrix[i][j], _ = strconv.Atoi(mas[j])
		}

	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i][j-1] + sum[i-1][j] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	for i := 0; i < z; i++ {
		value, _ = reader.ReadString('\n')
		value = strings.Replace(value, "\n", "", -1)
		mas = strings.Split(value, " ")
		x1, _ := strconv.Atoi(mas[0])
		y1, _ := strconv.Atoi(mas[1])
		x2, _ := strconv.Atoi(mas[2])
		y2, _ := strconv.Atoi(mas[3])
		s := sum[x2][y2] - sum[x1-1][y2] - sum[x2][y1-1] + sum[x1-1][y1-1]
		masS = append(masS, s)

	}
	for i, _ := range masS {
		fmt.Println(masS[i])
	}
}
