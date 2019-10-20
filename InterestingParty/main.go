package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	a, b := nextLine(), nextLine()
	f := strings.Split(a, " ")
	s := strings.Split(b, " ")
	myMain(f, s)
}

func myMain(first []string, second []string) {
	var counts map[string]int = map[string]int{}
	for i := 0; i < len(first); i++ {
		f := first[i]
		s := second[i]
		fv, fok := counts[f]
		if fok {
			counts[f] = fv + 1
		} else {
			counts[f] = 1
		}
		sv, sok := counts[s]
		if sok {
			counts[s] = sv + 1
		} else {
			counts[s] = 1
		}
	}
	res := 0
	for _, v := range counts {
		if v > res {
			res = v
		}
	}
	fmt.Println(res)
}
