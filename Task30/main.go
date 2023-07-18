package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, m, tempMax, temp int
	var result string

	fmt.Scanln(&n)
	matrix := make([][]int, n+1)
	matrixS := make([][]string, n+1)
	matrix[0] = append(matrix[0], 0)
	matrixS[0] = append(matrixS[0], "")
	for i := 1; i <= n; i++ {
		fmt.Scan(&temp)
		matrix[i] = append(matrix[i], temp)
		matrixS[i] = append(matrixS[i], strconv.Itoa(temp))
	}
	fmt.Scanln(&m)

	for i := 1; i <= m; i++ {
		fmt.Scan(&temp)
		matrix[0] = append(matrix[0], temp)
		matrixS[0] = append(matrixS[0], strconv.Itoa(temp))
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if matrix[i][0] == matrix[0][j] {
				if i == 1 || j == 1 {
					tempMax = 1
					result = matrixS[i][0]

				} else {
					tempMax = matrix[i-1][j-1] + 1
					result = matrixS[i-1][j-1] + " " + matrixS[i][0]
				}

				matrix[i] = append(matrix[i], tempMax)
				matrixS[i] = append(matrixS[i], result)
			} else {
				if i == 1 || j == 1 {
					if i == 1 && j == 1 {
						matrix[i] = append(matrix[i], 0)
						matrixS[i] = append(matrixS[i], "")
					} else {
						if i == 1 && j > 1 {
							matrix[i] = append(matrix[i], matrix[i][j-1])
							matrixS[i] = append(matrixS[i], matrixS[i][j-1])
						}
					}
					if j == 1 && i > 1 {
						matrix[i] = append(matrix[i], matrix[i-1][j])
						matrixS[i] = append(matrixS[i], matrixS[i-1][j])
					}

				} else {
					if matrix[i][j-1] > matrix[i-1][j] {
						tempMax = matrix[i][j-1]
						result = matrixS[i][j-1]
					} else {
						tempMax = matrix[i-1][j]
						result = matrixS[i-1][j]
					}
					matrix[i] = append(matrix[i], tempMax)
					matrixS[i] = append(matrixS[i], result)
				}
			}
		}
	}
	/*for i, _ := range matrix {
		fmt.Println(matrix[i])
	}
	fmt.Println()
	for i, _ := range matrix {
		fmt.Println(matrixS[i])
	}*/
	fmt.Println(matrixS[n][m])

}
