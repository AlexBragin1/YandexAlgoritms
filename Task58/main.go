package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, k, count, num int
	var simbol byte
	var name string
	fileOpen, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileOpen.Close()
	text := bufio.NewReader(fileOpen)
	fmt.Fscanln(text, &n, &m)
	fmt.Println(n, m)
	queue := make([]byte, 0, m*n)
	stackSimbol := make([]byte, 0, m*n)

	for i := 0; i < n; i++ {

		for j := 0; j < m; j++ {
			num += j
			fmt.Fscan(text, &queue[num])
			//fmt.Println(simbol)
			//queue = append(queue, simbol)
		}
		stackSimbol = append(stackSimbol, 32)
		fmt.Fscan(text)
	}
	fmt.Fscanln(text, &k)
	fmt.Println(queue)
	num = 1
	fmt.Println(queue)
	for i := 1; i <= k; i++ {
		fmt.Fscanln(text, &name, &count, &simbol)
		for j := 0; j < count; j++ {
			num += j
			stackSimbol[num] = simbol
		}
	}
	fmt.Println(stackSimbol)
	fileWrite, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileWrite.Close()
	result := bufio.NewWriter(fileWrite)
	defer result.Flush()
	for i := 0; i < n; i++ {
		for j := 0; j < m+1; j++ {
			if queue[j] == 32 {
				fmt.Fprint(result, stackSimbol[len(stackSimbol)-i])
			}
		}
		fmt.Fprintln(result)
	}

}
