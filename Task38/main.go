package main

import (
	"bufio"
	"fmt"
	"os"
)

type myGraph map[int][]int

func (g *myGraph) initGraph() {

}
func main() {
	var m, n, s, t, q int
	dx := [8]int{}
	dy := [8]int{}
	fileOpen, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileOpen.Close()
	text := bufio.NewReader(fileOpen)
	fmt.Fscanln(text, &n, &m, &s, &t, &q)
	initGrapph
}
