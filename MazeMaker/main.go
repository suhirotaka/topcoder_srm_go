package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveMaze(maze [][]byte, startRow int, startCol int, moveRow []int, moveCol []int) int {
	n := len(maze)
	m := len(maze[0])
	var q [][3]int
	g := make([][]int, n)
	for i, _ := range g {
		g[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			g[i][j] = -1
		}
	}
	q = append(q, [3]int{startRow, startCol, 0})
	g[startRow][startCol] = 0
	for len(q) > 0 {
		// fmt.Println(q)
		pos := q[0]
		q = q[1:]
		for i := 0; i < len(moveRow); i++ {
			row := pos[0] + moveRow[i]
			col := pos[1] + moveCol[i]
			if row >= 0 && row < n && col >= 0 && col < m {
				if g[row][col] == -1 && string(maze[row][col]) == "." {
					q = append(q, [3]int{row, col, (pos[2] + 1)})
					g[row][col] = pos[2] + 1
				}
			}
		}
	}
	// fmt.Println(moveRow, moveCol)

	fmt.Println(g)
	res := -1
	for i, g1 := range g {
		for j, g2 := range g1 {
			if g2 == -1 && string(maze[i][j]) == "." {
				return -1
			}
			if g2 > res {
				res = g2
			}
		}
	}
	return res
}

func myMain(maze [][]byte, startRow int, startCol int, moveRow []int, moveCol []int) {
	res := solveMaze(maze, startRow, startCol, moveRow, moveCol)
	fmt.Println(res)
}

func main() {
	var n int
	fmt.Scanln(&n)
	var maze [][]byte
	for i := 0; i < n; i++ {
		var r []byte
		fmt.Scanln(&r)
		maze = append(maze, r)
	}

	var startRow, startCol int
	fmt.Scanln(&startRow)
	fmt.Scanln(&startCol)

	var moveRow, moveCol []int
	var moveRowI, moveColI string
	fmt.Scanln(&moveRowI)
	moveRowS := strings.Split(moveRowI, ",")
	for _, v := range moveRowS {
		vv, _ := strconv.Atoi(v)
		moveRow = append(moveRow, vv)
	}
	fmt.Scanln(&moveColI)
	for _, v := range strings.Split(moveColI, ",") {
		vv, _ := strconv.Atoi(v)
		moveCol = append(moveCol, vv)
	}
	myMain(maze, startRow, startCol, moveRow, moveCol)
}
