package main

import (
	"fmt"
)

type stack struct {
	stk []int
}

func (s *stack) push(n int) {

	s.stk = append(s.stk, n)

}
func (s *stack) pop() {

	if s.size() == 0 {
		fmt.Println("error")
	} else {

		s.stk = s.stk[:len(s.stk)-1]
	}

}
func (s *stack) size() int {
	return len(s.stk)
}
func (s *stack) back() int {
	if s.size() == 0 {
		return 0
	} else {
		return s.stk[len(s.stk)-1]
	}
}

func main() {

	var n, numVagon int

	var vagon int

	mystack1 := stack{}
	tupik := stack{}
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vagon)

		mystack1.stk = append(mystack1.stk, vagon)
	}
	for i, j := 0, mystack1.size()-1; i < j; i, j = i+1, j-1 {
		mystack1.stk[i], mystack1.stk[j] = mystack1.stk[j], mystack1.stk[i]
	}
	numVagon = 1

	for numVagon < n {
		if numVagon != tupik.back() && mystack1.size() == 0 {
			break
		}
		if tupik.back() == numVagon {
			tupik.pop()
			numVagon++

		} else {
			tupik.push(mystack1.back())
			mystack1.pop()
		}

	}
	if numVagon == n {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
