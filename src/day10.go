package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() [][]int {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var grid [][]int
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		var vals []int
		for _, dig := range row {
			val, _ := strconv.Atoi(dig)
			vals = append(vals, val)
		}
		grid = append(grid, vals)
	}
	return grid
}

type tuple struct {
	x int
	y int
}

const START = 0
const END = 9

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func main() {
	grid := parseInput()
	n, m, ans := len(grid), len(grid[0]), 0

	unique_vis := make(map[tuple]bool)
	cnt_trailhead := 0

	var dfs func(x, y int)
	var chk func(x, y int) bool

	chk = func(x, y int) bool {
		return x >= 0 && y >= 0 && x < n && y < m
	}

	dfs = func(x, y int) {
		if grid[x][y] == END {
			cnt_trailhead++
			unique_vis[tuple{x, y}] = true
			return
		}

		for d := range dx {
			nx, ny := x+dx[d], y+dy[d]
			if chk(nx, ny) && grid[nx][ny] == grid[x][y]+1 {
				dfs(nx, ny)
			}
		}
	}

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == START {
				dfs(x, y)
				ans += len(unique_vis)
				unique_vis = make(map[tuple]bool)
			}
		}
	}

	fmt.Println(ans, cnt_trailhead)
}
