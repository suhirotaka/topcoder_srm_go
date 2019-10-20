package main

import (
	"fmt"
	"math"
	"strconv"
)

func myMain(base int) {
	var res []int
	maxTry := int(math.Pow(float64(base), float64(3))) - 1
	for m := 2; m < base; m++ {
		good := true
		for n := 1; n <= maxTry; n++ {
			pi := strconv.FormatInt(int64(m*n), base)
			ps := []byte(pi)
			sum := 0
			for _, p := range ps {
				pp, _ := strconv.ParseInt(string(p), base, 64)
				sum += int(pp)
			}
			if sum%m != 0 {
				good = false
			}
		}
		if good {
			res = append(res, m)
		}
	}
	fmt.Println(res)
}

func main() {
	var n int
	fmt.Scan(&n)
	myMain(n)
}
