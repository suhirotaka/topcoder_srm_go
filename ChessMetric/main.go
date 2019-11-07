package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var ms [51][100][100]int
	var size, numMoves int
	var start, finish string
	var startX, startY, finishX, finishY int

	fmt.Scanln(&size)
	fmt.Scanln(&start)
	fmt.Scanln(&finish)
	fmt.Scanln(&numMoves)

	starts := strings.Split(start, ",")
	finishes := strings.Split(finish, ",")
	startX, _ = strconv.Atoi(starts[0])
	startY, _ = strconv.Atoi(starts[1])
	finishX, _ = strconv.Atoi(finishes[0])
	finishY, _ = strconv.Atoi(finishes[1])

	ms[0][startX][startY] = 1

	for n := 1; n <= numMoves; n++ {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				var bps [][2]int
				bps = append(bps, [2]int{i - 2, j + 1})
				bps = append(bps, [2]int{i - 2, j - 1})

				bps = append(bps, [2]int{i - 1, j + 2})
				bps = append(bps, [2]int{i - 1, j + 1})
				bps = append(bps, [2]int{i - 1, j})
				bps = append(bps, [2]int{i - 1, j - 1})
				bps = append(bps, [2]int{i - 1, j - 2})

				bps = append(bps, [2]int{i, j + 1})
				bps = append(bps, [2]int{i, j - 1})

				bps = append(bps, [2]int{i + 1, j + 2})
				bps = append(bps, [2]int{i + 1, j + 1})
				bps = append(bps, [2]int{i + 1, j})
				bps = append(bps, [2]int{i + 1, j - 1})
				bps = append(bps, [2]int{i + 1, j - 2})

				bps = append(bps, [2]int{i + 2, j + 1})
				bps = append(bps, [2]int{i + 2, j - 1})

				a := 0
				for _, bp := range bps {
					if bp[0] >= 0 && bp[0] < size && bp[1] >= 0 && bp[1] < size {
						a += ms[n-1][bp[0]][bp[1]]
					}
				}
				ms[n][i][j] = a
			}
		}
	}

	fmt.Println(ms[numMoves][finishX][finishY])
}
