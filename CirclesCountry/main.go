package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	xs := strings.Split(sc.Text(), ",")
	sc.Scan()
	ys := strings.Split(sc.Text(), ",")
	sc.Scan()
	rs := strings.Split(sc.Text(), ",")
	sc.Scan()
	x1, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	y1, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	x2, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	y2, _ := strconv.Atoi(sc.Text())

	n := len(xs)
	c := 0
	for i := 0; i < n; i++ {
		x, _ := strconv.Atoi(xs[i])
		y, _ := strconv.Atoi(ys[i])
		r, _ := strconv.Atoi(rs[i])
		included1 := math.Pow(float64(x-x1), 2)+math.Pow(float64(y-y1), 2) < math.Pow(float64(r), 2)
		included2 := math.Pow(float64(x-x2), 2)+math.Pow(float64(y-y2), 2) < math.Pow(float64(r), 2)
		if included1 != included2 {
			c++
		}
	}
	fmt.Println(c)
}
