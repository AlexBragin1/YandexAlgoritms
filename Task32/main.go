package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type myGraph map[int][]int
type visited map[int]int
type sv struct {
	kol  int
	comp string
}

func (g *myGraph) dfs(now int, v *visited, marker int, p *sv) {

	(*v)[now] = marker
	p.kol++
	p.comp += strconv.Itoa(now) + " "

	for _, ngis := range (*g)[now] {
		if ngis > 0 && (*v)[ngis] == 0 {
			g.dfs(ngis, v, marker, p)
		}
	}

}

func main() {
	var m, n, e1, e2 int
	var res int

	fileOpen, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fileOpen.Close()
	text := bufio.NewReader(fileOpen)

	fmt.Fscanln(text, &m, &n)
	graph := make(myGraph, m+1)
	visa := make(visited, m+1)
	marker := make([]sv, m+1)
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

	for i := 1; i <= m; i++ {

		if visa[i] == 0 {
			graph.dfs(i, &visa, res+1, &marker[res])
			res++
		}

	}

	fmt.Fprintln(result, res)
	for i := 0; i < res; i++ {
		fmt.Fprintln(result, marker[i].kol)
		fmt.Fprintln(result, marker[i].comp)
	}
}
