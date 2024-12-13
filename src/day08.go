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

type cell struct {
	x int
	y int
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

func getAntennas(grid []string) map[byte][]cell {
	antennas := make(map[byte][]cell)
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] != '.' {
				antennas[grid[x][y]] = append(antennas[grid[x][y]], cell{x, y})
			}
		}
	}
	return antennas
}

func isCollinear(x1, y1, x2, y2, x3, y3 int) bool {
	dot := x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)
	return (dot == 0)
}

func getAntiNodesCnt(grid []string, antennas map[byte][]cell, onlyT bool) int {
	cnt := 0
	for x := range grid {
		for y := range grid[x] {
			antinode := false
			for _, locs := range antennas {
				for a := 0; a < len(locs); a++ {
					if antinode {
						break
					}
					for b := a + 1; b < len(locs); b++ {
						if antinode {
							break
						}
						x1, y1 := locs[a].x, locs[a].y
						x2, y2 := locs[b].x, locs[b].y
						d1 := (x-x1)*(x-x1) + (y-y1)*(y-y1)
						d2 := (x-x2)*(x-x2) + (y-y2)*(y-y2)
						if !onlyT && (d1 != 4*d2 && d2 != 4*d1) {
							continue
						}
						if isCollinear(x, y, x1, y1, x2, y2) {
							antinode = true
						}
					}
				}
			}
			if antinode {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	grid := parseInput()
	antennas := getAntennas(grid)
	fmt.Println(getAntiNodesCnt(grid, antennas, false))
	fmt.Println(getAntiNodesCnt(grid, antennas, true))

}
