package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var ms [50]int
	ms[0] = 1
	for i := 2; i <= n; i += 2 {
		sum := 0
		for j := 0; j <= i-2; j += 2 {
			sum += ms[j] * ms[i-j-2]
		}
		ms[i] = sum
	}
	fmt.Println(ms[n])
}
