package main

import (
	"bufio"
	"fmt"
	"os"
)

type myGraph map[int][]int
type visited map[int]int

func (g *myGraph) dfs(now int, v *visited, q *[]int) bool {
	if (*v)[now] == 1 {
		return true
	}
	if (*v)[now] == 2 {
		return false
	}
	(*v)[now] = 1
	*q = append(*q, now)
	for _, ngis := range (*g)[now] {
		if ngis > 0 && (*v)[ngis] == 0 {

			return true
		}
	}
	(*v)[now] = 2

	return false
}

func main() {
	var n, a int
	var r string
	var mas []int
	fileOpen, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileOpen.Close()
	text := bufio.NewReader(fileOpen)
	fmt.Fscanln(text, &n)

	graph := make(myGraph, n+1)
	visa := make(visited, n+1)
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
	r = "NO"
	for i := 1; i <= n; i++ {
		if visa[i] == 0 {
			if graph.dfs(i, &visa, &mas) {
				r = "YES"
				break
			} else {
				continue
			}
		}
	}
	fmt.Println(graph)
	fmt.Println(r)
	fmt.Println(mas)
}
