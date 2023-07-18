package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a int
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
	fmt.Fscanln(text, &n)
	arrN := make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(text, &a)
		arrN = append(arrN, a)
	}
	for i := 0; i < n; i++ {

	}
}
