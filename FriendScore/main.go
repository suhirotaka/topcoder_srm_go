package main

import (
	"fmt"
	"io"
)

func myMain(g [][]byte, n int) {
	rel := make([][]bool, n)
	for i := range rel {
		rel[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if string(g[i][j]) == "Y" {
				rel[i][j] = true
				for k, v := range g[j] {
					if string(v) == "Y" && i != k {
						rel[i][k] = true
					}
				}
			}
		}
	}

	maxSum := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if rel[i][j] {
				sum += 1
			}
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	fmt.Println(maxSum)
}

func main() {
	var mat [][]byte
	var l string
	n := 0
	for {
		_, err := fmt.Scanln(&l)
		if err == io.EOF {
			break
		}
		s := []byte(l)
		mat = append(mat, s)
		n += 1
	}
	myMain(mat, n)
}
