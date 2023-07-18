package main

import (
	"fmt"
)

type heaps struct {
	heap []int
	size int
}

func (h *heaps) heapify(indx int) {
	left := 2 * indx
	right := 2*indx + 1
	largest := indx

	if left < h.size && h.heap[left] > h.heap[largest] {
		largest = left
	}
	if right < h.size && h.heap[right] > h.heap[largest] {
		largest = right
	}
	if largest != indx {
		h.heap[indx], h.heap[largest] = h.heap[largest], h.heap[indx]
		h.heapify(largest)
	}
}

func (h *heaps) buildHeap() {
	h.size = len(h.heap)
	for i := (len(h.heap) - 1) / 2; i >= 0; i-- {
		h.heapify(i)
	}

}
func (h *heaps) sortPiramida() {
	h.buildHeap()
	for j := len(h.heap) - 1; j > 0; j-- {
		h.heap[0], h.heap[j] = h.heap[j], h.heap[0]
		h.size = h.size - 1

		h.heapify(0)

	}
}
func main() {

	var val, n int
	myheap := heaps{}
	fmt.Scanln(&n)
	for j := 1; j <= n; j++ {
		fmt.Scan(&val)
		myheap.heap = append(myheap.heap, val)
	}
	myheap.sortPiramida()
	for j, _ := range myheap.heap {
		fmt.Print(myheap.heap[j], " ")
	}
}
