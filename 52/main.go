package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int

	fileOpen, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileOpen.Close()
	text := bufio.NewReader(fileOpen)
	fileWrite, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileWrite.Close()
	result := bufio.NewWriter(fileWrite)
	defer result.Flush()
	fmt.Fscanln(text, &t)
	arr := make([]int, 0, t)
	for i := 0; i < t; i++ {

		fmt.Fscan(text, &n)
		arr = append(arr, n)

	}
	i := 0
	for len(arr) > 1 {
		if i == len(arr)-1 {
			arr = arr[:len(arr)-1]
			i = 0

		} else {
			arr[i] = arr[i] + arr[i+1]
			i++
		}
		fmt.Println(arr)
	}

	fmt.Fprintln(result, arr[0])
}
