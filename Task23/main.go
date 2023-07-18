package main

import (
	"fmt"
)

func main() {
	var n, num int
	fmt.Scan(&n)
	minOp := make([]int, 0, n)
	minOp = append(minOp, 0)
	for i := 2; i <= n; i++ {
		if i-1 > 0 {
			minOp = append(minOp, minOp[i-2]+1)
		}
		if i%2 == 0 {
			if minOp[i-1] > minOp[(i/2)-1]+1 {
				minOp[i-1] = minOp[(i/2)-1] + 1
			}
		}
		if i%3 == 0 {
			if minOp[i-1] > minOp[(i/3)-1]+1 {
				minOp[i-1] = minOp[(i/3)-1] + 1
			}
		}
	}
	masOp := make([]int, minOp[n-1]+1)
	masOp[0] = n
	j := 0
	num = n
	for i := minOp[n-1]; i >= 1; i-- {
		j++
		if num-1 > 0 {
			masOp[j] = num - 1
		}
		if num%2 == 0 {
			if minOp[masOp[j]] > minOp[(num/2)-1] {
				masOp[j] = num / 2
			}
		}
		if num%3 == 0 {
			if minOp[masOp[j]] > minOp[(num/3)-1] {
				masOp[j] = num / 3
			}
		}
		num = masOp[j]
	}
	fmt.Println(minOp[n-1])
	for i := len(masOp) - 1; i >= 0; i-- {
		fmt.Printf("%v ", masOp[i])
	}
}
