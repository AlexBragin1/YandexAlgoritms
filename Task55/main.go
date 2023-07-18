package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	res := make([]int, 3)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	value = strings.Trim(value, "\n")
	arr := strings.Split(value, " ")
	res[0], _ = strconv.Atoi(arr[0])
	res[1], _ = strconv.Atoi(arr[1])
	res[2], _ = strconv.Atoi(arr[2])
	sort.Ints(res)
	for i := 3; i < len(arr); i++ {
		now, _ := strconv.Atoi(arr[i])
		if res[0] < now {
			res[0] = now
			sort.Ints(res)
		}
	}
	fmt.Println(res[2], res[1], res[0])
}
