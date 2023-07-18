package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type prefix struct {
	val int
	ind int
}

func main() {

	var n, x, t, temp int
	fileOpen, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileOpen.Close()
	fileWrite, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileWrite.Close()
	result := bufio.NewWriter(fileWrite)
	defer result.Flush()
	text := bufio.NewReader(fileOpen)
	fmt.Fscanln(text, &n, &x, &t)
	arrN := make([]int, 0, n)
	prefixN := make([]prefix, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(text, &temp)
		arrN = append(arrN, temp)

		if arrN[i] == x {

			prefixN[i].val = 0
			prefixN[i].ind = 0
		}
		if arrN[i]-x > 0 {
			prefixN[i].val = arrN[i] - x
			prefixN[i].ind = i
		} else {

			prefixN[i].val = x - arrN[i]
			prefixN[i].ind = i

		}

	}

	sort.SliceStable(prefixN, func(i, j int) bool {
		return prefixN[i].val < prefixN[j].val
	})
	sum := 0
	ch := 0

	for i := 0; i < n; i++ {
		sum += prefixN[i].val
		if sum <= t {
			ch++
		} else {
			break
		}

	}
	fmt.Fprintln(result, ch)
	for i := 0; i < ch; i++ {
		fmt.Fprint(result, prefixN[i].ind+1, " ")
	}
}
