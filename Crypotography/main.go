package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func myMain(nums []int) {
	maxSum := 0
	for i, _ := range nums {
		prod := 1
		for j, n2 := range nums {
			if i == j {
				prod = prod * (n2 + 1)
			} else {
				prod = prod * n2
			}
		}
		if maxSum < prod {
			maxSum = prod
		}
	}
	fmt.Println(maxSum)
}

func main() {
	a := nextLine()
	ns := strings.Split(a, " ")
	var nums []int
	for _, v := range ns {
		v, _ := strconv.Atoi(v)
		nums = append(nums, v)
	}
	myMain(nums)
}
