package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type myGraph map[int][]int
type queue struct {
	stk []int
}

func (s *queue) push(n int) {

	s.stk = append(s.stk, n)

}
func (s *queue) pop() int {
	var value int
	if len(s.stk) == 0 {
		fmt.Println("error")
	} else {
		value = s.stk[0]
		s.stk = s.stk[1:len(s.stk)]
	}
	return value
}

func (g *myGraph) bfs(st int, en int, gu *queue) (int, string) {
	var res int
	var stroka string
	var resmas []int
	distance := make([]int, 0, len(*g)+1)
	path := make([]int, 0, len(*g)+1)
	for i := 1; i <= len(*g)+1; i++ {
		distance = append(distance, len(*g)+1)
		path = append(path, len(*g)+1)
	}

	distance[st] = 0
	gu.push(st)
	for len(gu.stk) > 0 {
		a := gu.pop()
		for _, ngis := range (*g)[a] {
			if ngis > 0 && distance[ngis] == len(*g)+1 {
				path[ngis] = a
				distance[ngis] = distance[a] + 1
				gu.push(ngis)
			}
		}
	}
	if distance[en] == len(*g)+1 {
		res = -1
	} else {
		res = distance[en]
		if res > 0 {
			for en != len(*g)+1 {
				resmas = append(resmas, en)
				en = path[en]

			}
			for i := len(resmas) - 1; i >= 0; i-- {

				stroka += strconv.Itoa(resmas[i]) + " "
			}
			stroka = "\n" + stroka
		}
	}
	fmt.Println(path)

	return res, stroka
}

func main() {
	var n, a, start, end int

	fileOpen, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileOpen.Close()
	text := bufio.NewReader(fileOpen)
	fmt.Fscanln(text, &n)

	graph := make(myGraph, n+1)

	myQueue := queue{}
	for i := 1; i <= n; i++ {
		graph[i] = append(graph[i], 0)

	}
	for i := 1; i <= n; i++ {
		for j := 1; j < n+1; j++ {
			fmt.Fscan(text, &a)

			if a == 1 && graph[i][0] == 0 {
				graph[i][0] = j
				continue
			}
			if a == 1 && graph[i][0] > 0 {
				graph[i] = append(graph[i], j)
			}

		}

	}
	fmt.Fscanln(text)
	fmt.Fscanln(text, &start, &end)
	fileWrite, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileWrite.Close()
	result := bufio.NewWriter(fileWrite)
	defer result.Flush()
	t1, t2 := graph.bfs(start, end, &myQueue)
	fmt.Fprintln(result, t1, t2)
}
