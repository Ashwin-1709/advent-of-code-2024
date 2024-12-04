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
		line := scanner.Text()
		grid = append(grid, line)
	}
	return grid
}

func ok(x, y, n, m int) bool {
	return (x >= 0) && (y >= 0) && (x < n) && (y < m)
}

func search(grid []string, x, y, dx, dy int) bool {
	n, m, step := len(grid), len(grid[0]), 0
	xmas := "XMAS"
	for {
		if step == len(xmas) {
			return true
		}
		if !ok(x, y, n, m) || grid[x][y] != xmas[step] {
			return false
		}
		x += dx
		y += dy
		step++
	}
}

func count(grid []string, x, y int) int {
	cnt := 0
	dx := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	dy := []int{0, -1, -1, -1, 0, 1, 1, 1}
	for i, _ := range dx {
		if search(grid, x, y, dx[i], dy[i]) {
			cnt++
		}
	}
	return cnt
}

func find(grid []string, x, y int) bool {
	n, m := len(grid), len(grid[x])
	if !ok(x-1, y-1, n, m) ||
		!ok(x+1, y-1, n, m) ||
		!ok(x+1, y+1, n, m) ||
		!ok(x-1, y+1, n, m) {
		return false
	}

	l := string(grid[x-1][y-1]) + string(grid[x+1][y+1])
	r := string(grid[x-1][y+1]) + string(grid[x+1][y-1])
	return (l == "SM" || l == "MS") &&
		(r == "SM" || r == "MS")
}

func main() {
	grid := parseInput()
	xmas, mas := 0, 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			switch grid[x][y] {
			case 'X':
				xmas += count(grid, x, y)
			case 'A':
				if find(grid, x, y) {
					mas++
				}
			default:
				continue
			}
		}
	}

	fmt.Println(xmas, mas)
}
