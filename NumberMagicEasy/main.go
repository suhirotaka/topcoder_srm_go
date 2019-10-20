package main

import (
	"fmt"
	"math"
)

func myMain(ans []byte, cs [16]int) {
	sum := 0
	for i, a := range ans {
		if string(a) == "Y" {
			sum += int(math.Pow(float64(2), float64(i)))
		}
	}
	fmt.Println(cs[sum])
}

func main() {
	var cs [16]int
	cs[0] = 16
	cs[1] = 8
	cs[2] = 12
	cs[3] = 4
	cs[4] = 14
	cs[5] = 6
	cs[6] = 10
	cs[7] = 2
	cs[8] = 15
	cs[9] = 7
	cs[10] = 11
	cs[11] = 3
	cs[12] = 13
	cs[13] = 5
	cs[14] = 9
	cs[15] = 1
	var ans []byte
	fmt.Scanln(&ans)
	myMain(ans, cs)
}
