package main

import "fmt"

func main() {
	var s string
	var i byte
	fmt.Scan(&s)
	m := make(map[byte]int, len(s))
	for i, _ := range s {
		m[s[i]] += (i + 1) * (len(s) - i)
	}

	for i = 97; i <= 122; i++ {
		if m[i] != 0 {
			fmt.Printf("%v: %v\n", string(i), m[i])
		}
	}
}
