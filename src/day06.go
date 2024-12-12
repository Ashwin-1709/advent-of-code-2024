package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInput() []string {
	f, err := os.Open("input.txt")
	handleError(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func IsCycle(g []string, stx, sty, tx, ty int) bool {
	if tx == stx && ty == sty {
		return false
	}
	turn, n, m := 0, len(g), len(g[0])
	seen := make([][][]bool, n)
	grid := make([]string, len(g))
	copy(grid, g)
	for i := range grid {
		seen[i] = make([][]bool, m)
		for j := range seen[i] {
			seen[i][j] = make([]bool, 4)
		}
	}
	row := grid[tx]
	grid[tx] = row[:ty] + "#" + row[ty+1:]

	for {
		if seen[stx][sty][turn%4] {
			return true
		}
		seen[stx][sty][turn%4] = true
		if stx == 0 || stx == n-1 || sty == 0 || sty == m-1 {
			break
		}
		nx, ny := stx+dx[turn%4], sty+dy[turn%4]
		cnt_turn := 0
		for {
			if grid[nx][ny] != '#' {
				break
			}
			if cnt_turn == 4 {
				return true
			}
			turn++
			cnt_turn++
			nx, ny = stx+dx[turn%4], sty+dy[turn%4]
		}
		stx, sty = nx, ny
	}
	return false
}

func walk(grid []string, stx, sty int) int {
	vis, turn, n, m := 0, 0, len(grid), len(grid[0])
	seen := make([][]bool, n)
	for i := range grid {
		seen[i] = make([]bool, m)
	}

	for {
		seen[stx][sty] = true
		if stx == 0 || stx == n-1 || sty == 0 || sty == m-1 {
			break
		}
		nx, ny := stx+dx[turn%4], sty+dy[turn%4]
		for {
			if grid[nx][ny] != '#' {
				break
			}
			turn++
			nx, ny = stx+dx[turn%4], sty+dy[turn%4]
		}
		stx, sty = nx, ny
	}

	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if seen[x][y] {
				vis++
			}
		}
	}
	return vis
}

func main() {
	grid := parseInput()
	trap := 0
	var stx, sty int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == '^' {
				stx, sty = x, y
			}
		}
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if IsCycle(grid, stx, sty, x, y) {
				trap++
			}
		}
	}

	fmt.Println(walk(grid, stx, sty), trap)
}
