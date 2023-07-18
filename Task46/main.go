package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a, k int
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
	keyb := make([]int, 0, n)
	keyboard := make(map[int]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(text, &a)
		keyb = append(keyb, a)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(text, &a)
		keyboard[keyb[i]] = a
	}
	fmt.Fscanln(text)
	fmt.Fscanln(text, &k)
	res := 0
	prewr := 0

	for i := 0; i < k; i++ {
		fmt.Fscan(text, &a)
		if i == 0 {
			prewr = keyboard[a]
			continue
		}

		if prewr != keyboard[a] && keyboard[a] > 0 {
			res++
			prewr = keyboard[a]
			continue
		}

	}
	fmt.Fprint(result, res)
}
