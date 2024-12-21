package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INPUT_FILE = "input.txt"
	ROBOT      = rune('@')
	WALL       = rune('#')
	DOT        = rune('.')
	BALL       = rune('O')
)

var (
	n = -1
	m = -1
)

func place(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func parseInput() ([]string, string) {
	f, _ := os.Open(INPUT_FILE)
	defer f.Close()
	var grid []string
	command := ""
	flag := true
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			flag = false
			continue
		}
		if flag {
			grid = append(grid, line)
		} else {
			command += line
		}
	}

	return grid, command
}

func print(grid []string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func ok(x, y int) bool {
	return x >= 1 && x < n-1 && y >= 1 && y < m-1
}

func pushBall(grid *[]string, wx, wy, dx, dy int) {
	nx, ny := wx+dx, wy+dy
	if !ok(nx, ny) {
		return
	}
	switch (*grid)[nx][ny] {
	case 'O':
		pushBall(grid, nx, ny, dx, dy)
		if (*grid)[nx][ny] == '.' {
			(*grid)[nx] = place((*grid)[nx], BALL, ny)
			(*grid)[wx] = place((*grid)[wx], DOT, wy)
		}
	case '.':
		(*grid)[nx] = place((*grid)[nx], BALL, ny)
		(*grid)[wx] = place((*grid)[wx], DOT, wy)
	default:
	}
}

func move(grid *[]string, rx, ry, dx, dy int) (int, int) {
	nx, ny := rx+dx, ry+dy
	if !ok(nx, ny) {
		return rx, ry
	}

	switch (*grid)[nx][ny] {
	case 'O':
		pushBall(grid, nx, ny, dx, dy)
		if (*grid)[nx][ny] == '.' {
			(*grid)[nx] = place((*grid)[nx], ROBOT, ny)
			(*grid)[rx] = place((*grid)[rx], DOT, ry)
			rx, ry = nx, ny
		}
	case '.':
		(*grid)[nx] = place((*grid)[nx], ROBOT, ny)
		(*grid)[rx] = place((*grid)[rx], DOT, ry)
		rx, ry = nx, ny
	default:
	}
	return rx, ry
}

func main() {
	grid, command := parseInput()
	n, m = len(grid), len(grid[0])

	rbx, rby := -1, -1
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '@' {
				rbx, rby = x, y
			}
		}
	}

	for i := range command {
		switch command[i] {
		case '>':
			rbx, rby = move(&grid, rbx, rby, 0, 1)
		case '<':
			rbx, rby = move(&grid, rbx, rby, 0, -1)
		case '^':
			rbx, rby = move(&grid, rbx, rby, -1, 0)
		default:
			rbx, rby = move(&grid, rbx, rby, 1, 0)
		}
		fmt.Println(i, "/", len(command))
	}

	ans := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'O' {
				ans += 100*x + y
			}
		}
	}

	fmt.Println(ans)
}
