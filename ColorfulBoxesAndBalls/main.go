package main

import "fmt"

func main() {
	var numRed, numBlue, onlyRed, onlyBlue, bothColors int
	fmt.Scanln(&numRed)
	fmt.Scanln(&numBlue)
	fmt.Scanln(&onlyRed)
	fmt.Scanln(&onlyBlue)
	fmt.Scanln(&bothColors)

	max := onlyRed*numRed + onlyBlue*numBlue
	count := 0
	for {
		if count == numRed || count == numBlue {
			break
		}
		t := max - onlyRed - onlyBlue + bothColors*2
		if t > max {
			max = t
		}
		count += 1
	}
	fmt.Println(max)
}
