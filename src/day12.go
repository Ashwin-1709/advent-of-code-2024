package main

import (
	"bufio"
	"fmt"
	"os"
)

type BorderInfo struct {
	x   int
	y   int
	dir int
}

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

func trace(border map[BorderInfo]bool) int {
	sides := 0

	check := func(curx, cury, id, loc int) {
		for {
			nx, ny := curx+dx[id], cury+dy[id]
			if !ok(nx, ny) {
				break
			}
			nxt := BorderInfo{x: nx, y: ny, dir: loc}
			_, found := border[nxt]
			if !found {
				break
			}
			border[nxt] = false
			curx, cury = nx, ny
		}
	}

	for corner, unseen := range border {
		if !unseen {
			continue
		}
		border[corner] = false
		sides++
		x, y, id := corner.x, corner.y, corner.dir
		if id < 2 {
			check(x, y, 2, id)
			check(x, y, 3, id)
		} else {
			check(x, y, 0, id)
			check(x, y, 1, id)
		}
	}
	return sides
}

func main() {
	grid := parseInput()
	n, m = len(grid), len(grid[0])
	vis := make([][]bool, n)
	for x := range grid {
		vis[x] = make([]bool, m)
	}

	part_1, part_2 := 0, 0
	for x := range grid {
		for y := range grid[x] {
			if vis[x][y] {
				continue
			}

			perimeter, area := 0, 0
			border := make(map[BorderInfo]bool)

			var dfs func(x, y int)
			dfs = func(x, y int) {
				vis[x][y] = true
				area++

				for i := range dx {
					nx, ny := x+dx[i], y+dy[i]
					if !ok(nx, ny) || (ok(nx, ny) && grid[nx][ny] != grid[x][y]) {
						perimeter++
						border[BorderInfo{x: x, y: y, dir: i}] = true
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
			part_2 += trace(border) * area
		}
	}

	fmt.Println(part_1, part_2)
}
