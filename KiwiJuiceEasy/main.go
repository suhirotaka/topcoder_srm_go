package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func myMain(capacities []int, bottles []int, fromId []int, toId []int) {
	for i := 0; i < len(fromId); i++ {
		f := fromId[i]
		t := toId[i]
		var v int
		if capacities[t] < bottles[t]+bottles[f] {
			v = capacities[t] - bottles[t]
		} else {
			v = bottles[f]
		}
		bottles[f] -= v
		bottles[t] += v
	}
	fmt.Printf("%v", bottles)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	a, b, c, d := nextLine(), nextLine(), nextLine(), nextLine()
	cas := strings.Split(a, " ")
	bos := strings.Split(b, " ")
	frs := strings.Split(c, " ")
	tos := strings.Split(d, " ")
	var capacities, bottles, fromId, toId []int
	for _, ca := range cas {
		cac, _ := strconv.Atoi(ca)
		capacities = append(capacities, cac)
	}
	for _, bo := range bos {
		boc, _ := strconv.Atoi(bo)
		bottles = append(bottles, boc)
	}
	for _, fr := range frs {
		frc, _ := strconv.Atoi(fr)
		fromId = append(fromId, frc)
	}
	for _, to := range tos {
		toc, _ := strconv.Atoi(to)
		toId = append(toId, toc)
	}
	myMain(capacities, bottles, fromId, toId)
	//	fmt.Printf("%v", capacities)
	//	fmt.Printf("%v", bottles)
	//	fmt.Printf("%v", fromId)
	//	fmt.Printf("%v", toId)
}
