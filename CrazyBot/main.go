package main

import "fmt"

func myMain(dists [4]int, locs [50][50]bool, locx int, locy int, n int, prob float64) float64 {
	if locs[locx+15][locy+15] {
		return 0.0
	}
	if n == 0 {
		return prob
	}
	locs[locx+15][locy+15] = true
	p1 := myMain(dists, locs, locx+1, locy, n-1, prob*float64(dists[0])*0.01)
	p2 := myMain(dists, locs, locx-1, locy, n-1, prob*float64(dists[1])*0.01)
	p3 := myMain(dists, locs, locx, locy-1, n-1, prob*float64(dists[2])*0.01)
	p4 := myMain(dists, locs, locx, locy+1, n-1, prob*float64(dists[3])*0.01)
	return p1 + p2 + p3 + p4
}

func main() {
	var n, east, west, south, north int
	fmt.Scanln(&n)
	fmt.Scanln(&east)
	fmt.Scanln(&west)
	fmt.Scanln(&south)
	fmt.Scanln(&north)
	dists := [4]int{east, west, south, north}
	var locs [50][50]bool
	r := myMain(dists, locs, 0, 0, n, 1.0)
	fmt.Printf("%f", r)
}
