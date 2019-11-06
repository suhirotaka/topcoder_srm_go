package main

import (
	"fmt"
	"io"
)

var ms [50]int

func calcSalary(ss [][]byte, n int) int {
	s := ss[n]
	salary := 0
	for i, c := range s {
		if string(c) == "Y" {
			if ms[i] > 0 {
				// fmt.Println("get from memo")
				salary += ms[i]
			} else {
				// fmt.Println("do recursion")
				i_salary := calcSalary(ss, i)
				salary += i_salary
			}
		}
	}
	if salary == 0 {
		salary = 1
	}
	ms[n] = salary
	return salary
}

func main() {
	var ss [][]byte
	for {
		var s []byte
		_, err := fmt.Scanln(&s)
		if err == io.EOF {
			break
		}
		ss = append(ss, s)
	}

	for i := 0; i < len(ss); i++ {
		calcSalary(ss, i)
	}

	ans := 0
	for _, m := range ms {
		ans += m
	}
	fmt.Println(ans)
}
