package main

import (
	"fmt"
	"strconv"
)

type myGraph map[int][]int
type visited map[]bool

func (g *myGraph) dfs(now int, v *visited, servSearch int) bool {

	flag := false
	(*v)[now] = true
	for _, ngis := range (*g)[now] {
		if ngis == servSearch {
			flag = true
			break
		} else if ngis > 0 && !(*v)[ngis] {
			flag = g.dfs(ngis, v, servSearch)
			if flag {
				break
			}
		}
	}
	return flag
}
func main() {
	var n, e1, e2, m, k, searchServer, temp int
	var res []string
	fmt.Scanln(&n)
	serv := make(myGraph, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scanln(&e1, &e2)
		serv[e1] = append(serv[e1], e2)
		serv[e2] = append(serv[e2], e1)
	}

	fmt.Scanln(&m)
	for i := 0; i < m; i++ {

		fmt.Scanln(&searchServer, &k)
		res = append(res, "")
		ch := 0
		for j := 0; j < k; j++ {
			fmt.Scan(&temp)

			visa := make(visited, n+1)

			if serv.dfs(temp, &visa, searchServer) {

				res[i] = res[i] + " " + strconv.Itoa(temp)
				ch++

			}
		}

		if ch == 0 {
			res[i] = "0"
		} else {
			res[i] = strconv.Itoa(ch) + res[i]
		}

	}
	for i := 0; i < m; i++ {
		fmt.Println(res[i])
	}
}
