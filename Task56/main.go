package main

import (
	"fmt"
	"sort"
)

func main() {
	var m, n int
	var color int64
	var res []int64
	var resA []int64
	var resB []int64
	fmt.Scanln(&m, &n)
	anya := make(map[int64]int, m)
	boris := make([]int64, n)
	for i := 0; i < m; i++ {
		fmt.Scanln(&color)
		anya[color]++
	}
	for i := 0; i < n; i++ {
		fmt.Scanln(&boris[i])
		if anya[boris[i]] > 0 {
			res = append(res, boris[i])
		} else {
			resB = append(resB, boris[i])
		}
	}
	sort.SliceIsSorted(res,
		func(i, j int) bool { return i < j})
	sort.SliceIsSorted(resB,
		func(i, j int) bool { return i < j})
	fmt.Println(len(res))
	for i:=0;i<len(res);i++{
		fmt.Print(res[i]," ")
	}
	for i,_:=range anya{
		if  any[res]
	}
	fmt.Println(len(resA))
	fmt.Println(len(resB))
	for i:=0;i<len(resB);i++{

	}
}
