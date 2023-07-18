package main

import (
	"fmt"
)

type heaps struct {
	heap []int
}

func (h *heaps) size() int {
	return len(h.heap)
}
func (h *heaps) heapify(i int) {
	left := 2 * i
	right := 2*i + 1
	largest := i

	if left < h.size() && h.heap[left] > h.heap[largest] {
		largest = left
	}
	if right < h.size() && h.heap[right] > h.heap[largest] {
		largest = right
	}
	if largest != i {
		h.heap[i], h.heap[largest] = h.heap[largest], h.heap[i]
		h.heapify(largest)
	}
}
func (h *heaps) insert(value int) {
	h.heap = append(h.heap, value)
	if h.size() != 1 {
		i := h.size() - 1

		for i > 0 && h.heap[i/2] < h.heap[i] {
			h.heap[i/2], h.heap[i] = h.heap[i], h.heap[i/2]
			i = i / 2
		}
	}
}
func (h *heaps) extract() int {
	if h.size() == 0 {
		fmt.Println("heap is epmty")
		return -1
	}
	max := h.heap[0]
	h.heap[0] = h.heap[h.size()-1]
	h.heap = h.heap[0 : h.size()-1]
	if h.size() > 1 {
		h.heapify(0)
	}
	return max
}
func main() {
	//var s string
	var val, n, oper int
	var res []int
	myheap := heaps{}

	fmt.Scanln(&n)

	for j := 1; j <= n; j++ {
		fmt.Scan(&oper)
		if oper == 0 {
			fmt.Scanln(&val)
			myheap.insert(val)
		}
		if oper == 1 {

			res = append(res, myheap.extract())

		}

	}
	for j, _ := range res {
		fmt.Println(res[j])
	}
}
