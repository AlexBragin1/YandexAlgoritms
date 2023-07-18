package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a2, res int
	var a1, a3 string
	var stackTovar []string
	var stackkol []int
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
	fmt.Fscanln(text, &n)
	for j := 0; j < n; j++ {
		fmt.Fscan(text, &a1)
		if a1 == "add" {
			fmt.Fscanln(text, &a2, &a3)
		} else {
			if a1 == "get" {
				fmt.Fscanln(text, &a3)
			} else {
				fmt.Fscanln(text, &a2)
			}
		}
		switch a1 {
		case "add":
			stackTovar = append(stackTovar, a3)
			stackkol = append(stackkol, a2)
		case "get":
			res = 0
			for i := 0; i < len(stackkol); i++ {
				if stackTovar[i] == a3 {
					res += stackkol[i]
				}
			}
			fmt.Fprintln(result, res)
		case "delete":
			for a2 != 0 {
				if stackkol[len(stackkol)-1]-a2 < 0 {
					a2 = stackkol[len(stackkol)-1] - a2
					stackkol = stackkol[:len(stackkol)-1]
					stackTovar = stackTovar[:len(stackTovar)-1]
					a2 = 0 - a2

				} else {
					stackkol[len(stackkol)-1] = stackkol[len(stackkol)-1] - a2
					break
				}
			}
		}
	}
}
