package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	dn := []int{}
	dn = append(dn, 2)
	dn = append(dn, 4)
	dn = append(dn, 7)
	for i := 3; i <= n; i++ {
		dn = append(dn, dn[i-1]+dn[i-2]+dn[i-3])

	}
	fmt.Println(dn[n-1])
}
