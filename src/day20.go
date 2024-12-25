package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INPUT_FILE = "input.txt"
	INF        = math.MaxInt64
	MIN_SAVE   = 100
)

var (
	n  = -1
	m  = -1
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

type Pair struct {
	x int
	y int
}

func parseInput() ([]string, int, int, int, int) {
	f, _ := os.Open(INPUT_FILE)
	defer f.Close()
	var grid []string
	var sx, sy, ex, ey int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
		x := len(grid) - 1
		for y := range grid[x] {
			if grid[x][y] == 'S' {
				sx, sy = x, y
			}
			if grid[x][y] == 'E' {
				ex, ey = x, y
			}
		}
	}
	return grid, sx, sy, ex, ey

}

func ok(x, y int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	grid, sx, sy, ex, ey := parseInput()
	n, m = len(grid), len(grid[0])

	bfs := func(sx, sy int) [][]int {
		dis := make([][]int, n)
		for x := range dis {
			dis[x] = make([]int, m)
			for y := range dis[x] {
				dis[x][y] = -1
			}
		}

		Q := make(chan Pair, n*m)
		dis[sx][sy] = 0
		Q <- Pair{sx, sy}
		for len(Q) > 0 {
			cur := <-Q
			for i := range dx {
				nx, ny := cur.x+dx[i], cur.y+dy[i]
				if ok(nx, ny) && grid[nx][ny] != '#' && dis[nx][ny] == -1 {
					dis[nx][ny] = dis[cur.x][cur.y] + 1
					Q <- Pair{nx, ny}
				}
			}
		}
		return dis
	}

	source, sink := bfs(sx, sy), bfs(ex, ey)
	var good []Pair
	part_1, part_2 := 0, 0

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] != '#' {
				good = append(good, Pair{x, y})
				continue
			}
			var valid []Pair
			for i := range dx {
				nx, ny := x+dx[i], y+dy[i]
				if ok(nx, ny) && grid[nx][ny] != '#' {
					valid = append(valid, Pair{nx, ny})
				}
			}

			for i := 0; i < len(valid); i++ {
				for j := 0; j < i; j++ {
					v1, v2 := valid[i], valid[j]
					total := source[v1.x][v1.y] + sink[v1.x][v1.y]
					best := min(source[v1.x][v1.y]+sink[v2.x][v2.y]+2, source[v2.x][v2.y]+sink[v1.x][v1.y]+2)
					if total-best >= MIN_SAVE {
						part_1++
					}
				}
			}
		}
	}

	for i := 0; i < len(good); i++ {
		for j := 0; j < i; j++ {
			v1, v2 := good[i], good[j]
			manhattan := abs(v1.x-v2.x) + abs(v1.y-v2.y)
			if manhattan <= 20 {
				total := source[v1.x][v1.y] + sink[v1.x][v1.y]
				best := min(source[v1.x][v1.y]+manhattan+sink[v2.x][v2.y], source[v2.x][v2.y]+manhattan+sink[v1.x][v1.y])
				if total-best >= MIN_SAVE {
					part_2++
				}
			}
		}
	}

	fmt.Println(part_1, part_2)
}
