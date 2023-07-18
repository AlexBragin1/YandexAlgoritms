package main

import (
	"fmt"
	"strconv"
	"strings"
)

func roundInt(n int) int {
	if n%2 != 0 {
		return n + 1
	}
	return n
}
func timeToStr(n int, flag bool) string {
	var res string
	if flag == true && n >= 24 {
		n = n - 24
	}
	if n < 10 {
		res = "0"
	}
	res += strconv.Itoa(n)
	return res
}
func main() {
	var s string
	var d int
	arrSec := make([]int, 0, 3)
	for i := 0; i < 3; i++ {
		fmt.Scanln(&s)
		times := strings.Split(s, ":")
		hour, _ := strconv.Atoi(times[0])
		min, _ := strconv.Atoi(times[1])
		sec, _ := strconv.Atoi(times[2])
		arrSec = append(arrSec, 3600*hour+60*min+sec)
	}

	if arrSec[2] < arrSec[0] {
		d = roundInt((24*3600 + arrSec[2] - arrSec[0])) / 2
	} else {
		d = roundInt(arrSec[2]-arrSec[0]) / 2
	}
	arrSec[1] += d

	s = timeToStr(arrSec[1]/3600, true) + ":" + timeToStr((arrSec[1]%3600)/60, false) + ":" + timeToStr(arrSec[1]%3600%60, false)

	fmt.Println(s)
}
