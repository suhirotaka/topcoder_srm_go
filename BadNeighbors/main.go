package main

import (
	"fmt"
	"strconv"
	"strings"
)

var ms [40]int

func badNeighbors(ss []string) int {
	n := len(ss)
	num0, _ := strconv.Atoi(ss[0])
	ms[0] = num0
	if n > 1 {
		num1, _ := strconv.Atoi(ss[1])
		if num0 > num1 {
			ms[1] = num0
		} else {
			ms[1] = num1
		}
	}
	for i := 2; i < n; i++ {
		num, _ := strconv.Atoi(ss[i])
		r1 := num + ms[i-2]
		r2 := ms[i-1]
		if r1 > r2 {
			ms[i] = r1
		} else {
			ms[i] = r2
		}
	}
	return ms[n-1]
}

func main() {
	var s string
	fmt.Scanln(&s)
	ss := strings.Split(s, ",")
	r1 := badNeighbors(ss[1:])
	r2 := badNeighbors(ss[:len(ss)-1])
	if r1 > r2 {
		fmt.Println(r1)
	} else {
		fmt.Println(r2)
	}
}
