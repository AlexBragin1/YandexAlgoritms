package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileOpen, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileOpen.Close()
	text := bufio.NewReader(fileOpen)
	text = strings.Trim(text, "\n")
	arr := strings.Split(string(text), " ")
}
