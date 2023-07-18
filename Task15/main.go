package main

import (
	"fmt"
)

type stack struct {
	cost []int
	nums []int
}

func (s *stack) push(price int, num int) {
	s.cost = append(s.cost, price)
	s.nums = append(s.nums, num)
}
func (s *stack) pop() {
	s.cost = s.cost[:len(s.cost)-1]
	s.nums = s.nums[:len(s.nums)-1]
}
func (s *stack) backPrice() int {

	return s.cost[len(s.cost)-1]

}
func (s *stack) backNums() int {

	return s.nums[len(s.nums)-1]

}
func (s *stack) size() int {
	return len(s.cost)
}
func main() {
	var n int

	st := stack{}
	fmt.Scanln(&n)
	res := make([]int, n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		res[i] = -1
	}

	for i := 0; i < n; i++ {
		for st.size() > 0 && arr[i] < st.backPrice() {
			res[st.backNums()] = i
			st.pop()
		}
		st.push(arr[i], i)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%v ", res[i])
	}

}
