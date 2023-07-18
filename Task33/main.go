package main

import (
	"bufio"
	"fmt"
	"os"
)

type myGraph map[int][]int
type visited map[int]int

func (g *myGraph) dfs(now int, v *visited, color int) string {
	var reslt string
	(*v)[now] = color
	reslt = "YES"

	for _, ngis := range (*g)[now] {
		if (*v)[now] == (*v)[ngis] {
			return "NO"
		}
		if ngis > 0 && (*v)[ngis] == 0 {
			reslt = g.dfs(ngis, v, 3-color)
		}
	}
	return reslt
}

func main() {
	var m, n, e1, e2 int
	var r string

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

	}
	for i := 1; i <= n; i++ {
		fmt.Fscanln(text, &e1, &e2)
		if e1 != e2 {
			if graph[e1][0] == 0 {
				graph[e1][0] = e2
			} else {
				graph[e1] = append(graph[e1], e2)
			}
			if graph[e2][0] == 0 {
				graph[e2][0] = e1
			} else {
				graph[e2] = append(graph[e2], e1)
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
	for i := 1; i < m; i++ {
		if visa[i] == 0 {
			if graph.dfs(i, &visa, 1) == "NO" {
				r = "NO"
				break
			} else {
				r = "YES"
				continue
			}
		}
	}

	fmt.Fprintln(result, r)
}
