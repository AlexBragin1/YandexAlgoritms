package main

import (
	"fmt"
	"os"
	"strconv"
)

type queue struct {
	stk []int
}

func (s *queue) push(n int) {
	s.stk = append(s.stk, n)
}
func (s *queue) pop() int {

	elem := s.stk[0]
	s.stk = s.stk[1:len(s.stk)]
	return elem
}

func (s *queue) size() int {
	return len(s.stk)
}
func (s *queue) front() int {

	return s.stk[0]

}
func main() {
	var steps int
	var win string

	firstQueue := queue{}
	secondQueue := queue{}
	bytes, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}
	for i, _ := range bytes {
		if bytes[i] >= '0' && bytes[i] <= '9' {
			if firstQueue.size() > 4 {
				secondQueue.push(int(bytes[i] - '0'))
			} else {
				firstQueue.push(int(bytes[i] - '0'))
			}
		}
	}

	for {
		if steps == 1000000 {
			win = "botva"
			break
		}

		if firstQueue.size() == 0 {
			win = "second"
			break
		}
		if secondQueue.size() == 0 {
			win = "first"
			break
		}
		steps++
		if firstQueue.front() == 0 && secondQueue.front() == 9 {
			firstQueue.push(firstQueue.pop())
			firstQueue.push(secondQueue.pop())

			continue
		}
		if firstQueue.front() == 9 && secondQueue.front() == 0 {

			secondQueue.push(firstQueue.pop())
			secondQueue.push(secondQueue.pop())

			continue
		}
		if firstQueue.front() > secondQueue.front() {
			firstQueue.push(firstQueue.pop())
			firstQueue.push(secondQueue.pop())

			continue
		} else {
			secondQueue.push(firstQueue.pop())
			secondQueue.push(secondQueue.pop())

			continue
		}

	}
	file, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	if win != "botva" {
		win = win + " " + strconv.Itoa(steps)
	}
	file.WriteString(win)
}
