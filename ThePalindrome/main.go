package main

import "fmt"

func myMain(s string) {
	ss1 := []byte(s)
	n := len(ss1)
	var ss2 []byte
	for i := n - 1; i >= 0; i-- {
		ss2 = append(ss2, ss1[i])
	}
	var adds int
	for i := 0; i <= n; i++ {
		sl1 := ss1[i:]
		sl2 := ss2[:(n - i)]
		matched := true
		for j := 0; j < n-i; j++ {
			if sl1[j] != sl2[j] {
				matched = false
				break
			}
		}
		if matched {
			adds = i
			break
		}
	}
	fmt.Println(n + adds)
}

func main() {
	var s string
	fmt.Scanln(&s)
	myMain(s)
}
