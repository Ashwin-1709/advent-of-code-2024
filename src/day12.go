package main

import (
	"bufio"
	"fmt"
	"os"
)

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}
var n, m = 0, 0

func ok(x, y int) bool {
	return x >= 0 && y >= 0 && x < n && y < m
}

func parseInput() []string {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

func main() {
	grid := parseInput()
	n, m = len(grid), len(grid[0])
	vis := make([][]bool, n)
	for x := range grid {
		vis[x] = make([]bool, m)
	}

	part_1 := 0
	for x := range grid {
		for y := range grid[x] {
			if vis[x][y] {
				continue
			}
			perimeter, area := 0, 0

			var dfs func(x, y int)
			dfs = func(x, y int) {
				vis[x][y] = true
				area++
				for i := range dx {
					nx, ny := x+dx[i], y+dy[i]
					if !ok(nx, ny) || (ok(nx, ny) && grid[nx][ny] != grid[x][y]) {
						perimeter++
						continue
					}
					if vis[nx][ny] {
						continue
					}
					dfs(nx, ny)
				}
			}

			dfs(x, y)
			part_1 += perimeter * area
		}
	}

	fmt.Println(part_1)
}
