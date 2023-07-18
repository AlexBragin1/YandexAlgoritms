package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type stack struct {
	stc []int
}

func (s *stack) push(n int) {
	s.stc = append(s.stc, n)
}
func (s *stack) pop() int {
	var value int
	if len(s.stc) == 0 {
		fmt.Println("error")
	} else {
		value = s.stc[len(s.stc)-1]
		s.stc = s.stc[:len(s.stc)-1]
	}
	return value
}

func main() {
	myStack := stack{}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	line := make([]byte, 64)
	for {
		_, err1 := file.Read(line)
		if err1 == io.EOF { // если конец файла
			break // выходим из цикла
		}

		for i, _ := range line {
			if i != 0 && line[i-1] == ' ' && line[i] == ' ' {

				break
			}

			switch {
			case '0' <= line[i] && line[i] <= '9':

				myStack.push(int(line[i] - '0'))

			case line[i] == '+':
				a := myStack.pop()
				b := myStack.pop()
				myStack.push(b + a)

			case line[i] == '*':
				a := myStack.pop()
				b := myStack.pop()
				myStack.push(a * b)
			case line[i] == '-':
				a := myStack.pop()
				b := myStack.pop()
				myStack.push(b - a)

			}
		}
	}
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	buffer := bufio.NewWriter(f)
	_, err = buffer.WriteString(strconv.Itoa(myStack.stc[0]))
	if err != nil {
		log.Fatal(err)
	}

	// flush buffered data to the file
	if err = buffer.Flush(); err != nil {
		log.Fatal(err)
	}

}
