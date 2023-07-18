package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var di, dj, max1, max2 int
	reader := bufio.NewReader(os.Stdin)

	value, _ := reader.ReadString('\n')
	value = strings.Trim(value, "\n")
	arrStreet := strings.Split(value, " ")
	di = -1
	dj = len(arrStreet)
	j := len(arrStreet) - 1
	max1 = 0
	max2 = 0
	for i := 0; i <= len(arrStreet)-1; i++ {
		fmt.Println(arrStreet[i], arrStreet[j])
		if arrStreet[i] == "1" && di == -1 {
			di = i
		}
		if (arrStreet[j] == "1" && dj == -1) || (dj == len(arrStreet) && arrStreet[j] == "1") {
			dj = j
		}
		if arrStreet[i] == "2" && di != -1 {
			if i-di > max1 {
				max1 = i - di
			}
			di = -1
		}
		if arrStreet[j] == "2" && dj != -1 && dj != len(arrStreet) {
			if dj-j > max2 {
				max2 = dj - j
			}
			dj = -1
		}
		j--
		fmt.Println(di, dj, max1, max2)
	}
	if max2 > max1 {
		fmt.Println(max2)
	} else {
		fmt.Println(max1)
	}
}
