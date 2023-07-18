package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type myGraph [][]int
type visited []int

func (g *myGraph) dfs(now int, v *visited, q *[]int) bool {

	if (*v)[now] == 1 {
		return true
	}
	if (*v)[now] == 2 {
		return false
	}
	(*v)[now] = 1
	for _, ngis := range (*g)[now] {
		if ngis > 0 && g.dfs(ngis, v, q) {
			return true
		}
	}
	(*v)[now] = 2
	*q = append(*q, now)
	return false
}

func main() {
	var m, n, e1, e2 int
	var sortres string
	var fl bool
	myQueue := []int{}
	fileOpen, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileOpen.Close()
	text := bufio.NewReader(fileOpen)

	fmt.Fscanln(text, &m, &n)
	graph := make(myGraph, m+1)
	visa := make(visited, m+1)

	for i := 1; i <= m; i++ {
		graph[i] = append(graph[i], 0)
		visa[i] = 0
	}
	for i := 1; i <= n; i++ {
		fmt.Fscanln(text, &e1, &e2)
		if e1 != e2 {
			if graph[e1][0] == 0 {
				graph[e1][0] = e2
			} else {
				graph[e1] = append(graph[e1], e2)
			}

		} else {
			if graph[e1][0] == 0 {
				graph[e1][0] = e1
			} else {
				graph[e1] = append(graph[e1], e1)
			}
		}
	}

	fileWrite, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileWrite.Close()
	result := bufio.NewWriter(fileWrite)
	defer result.Flush()

	for i := 1; i <= m; i++ {
		if visa[i] != 2 {
			fl = graph.dfs(i, &visa, &myQueue)
			if fl == true {
				break
			}
			for len(myQueue) > 0 {
				a := myQueue[0]
				myQueue = myQueue[1:len(myQueue)]
				sortres = strconv.Itoa(a) + " " + sortres
			}

		}
	}

	if fl == true {
		fmt.Fprint(result, -1)
	} else {
		fmt.Fprint(result, sortres)
	}
}
