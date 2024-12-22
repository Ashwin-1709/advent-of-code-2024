package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type State struct {
	x    int
	y    int
	dir  int
	cost int
}

type hp []State

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].cost < h[j].cost }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(State)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

const (
	INPUT_FILE = "input.txt"
	INF        = math.MaxInt64
	TURN_COST  = 1000
)

var (
	n  = -1
	m  = -1
	DX = []int{-1, 0, 1, 0}
	DY = []int{0, 1, 0, -1}
)

func parseInput() ([]string, int, int, int, int) {
	f, _ := os.Open(INPUT_FILE)
	defer f.Close()
	var grid []string
	sx, sy, tx, ty := -1, -1, -1, -1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
		x := len(grid) - 1
		for y := range grid[x] {
			switch grid[x][y] {
			case 'S':
				sx, sy = x, y
			case 'E':
				tx, ty = x, y
			default:
			}
		}
	}
	return grid, sx, sy, tx, ty
}

func ok(x, y int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
}

func dijkstra(grid []string, sx, sy int, dirs []int) [][][]int {
	h := &hp{}
	heap.Init(h)
	for _, d := range dirs {
		heap.Push(h, State{x: sx, y: sy, dir: d, cost: 0})
	}

	dis := make([][][]int, n)
	best := make([][]int, n)
	for x := range grid {
		dis[x] = make([][]int, m)
		best[x] = make([]int, m)
		for y := range grid[x] {
			dis[x][y] = make([]int, 4)
			best[x][y] = INF
			for i := 0; i < 4; i++ {
				dis[x][y][i] = INF
			}
		}
	}

	for len(*h) > 0 {
		cur := h.Pop().(State)
		x, y, dir, cost := cur.x, cur.y, cur.dir, cur.cost
		if dis[x][y][dir] < cost {
			continue
		}
		check := func(nx, ny, cst, add, ndir int) {
			if !ok(nx, ny) || dis[nx][ny][ndir] < cst+add || grid[nx][ny] == '#' {
				return
			}
			dis[nx][ny][ndir] = cst + add
			heap.Push(h, State{x: nx, y: ny, dir: ndir, cost: cst + add})
		}

		dis[x][y][dir] = cost
		cdx, cdy := DX[dir], DY[dir]
		check(x+cdx, y+cdy, cost, 1, dir)
		for i := range DX {
			ndx, ndy := DX[i], DY[i]
			dot := cdx*ndx + cdy*ndy
			if dot != 0 {
				continue
			}
			check(x+ndx, y+ndy, cost, TURN_COST+1, i)
		}
	}

	return dis
}

func rev_dijkstra(grid []string, sx, sy int, dirs []int) [][][]int {
	h := &hp{}
	heap.Init(h)
	for _, d := range dirs {
		heap.Push(h, State{x: sx, y: sy, dir: d, cost: 0})
	}

	dis := make([][][]int, n)
	best := make([][]int, n)
	for x := range grid {
		dis[x] = make([][]int, m)
		best[x] = make([]int, m)
		for y := range grid[x] {
			dis[x][y] = make([]int, 4)
			best[x][y] = INF
			for i := 0; i < 4; i++ {
				dis[x][y][i] = INF
			}
		}
	}

	for len(*h) > 0 {
		cur := h.Pop().(State)
		x, y, dir, cost := cur.x, cur.y, cur.dir, cur.cost
		if dis[x][y][dir] < cost {
			continue
		}
		dis[x][y][dir] = cost
		check := func(nx, ny, cst, add, ndir int) {
			if !ok(nx, ny) || dis[nx][ny][ndir] < cst+add || grid[nx][ny] == '#' {
				return
			}
			dis[nx][ny][ndir] = cst + add
			heap.Push(h, State{x: nx, y: ny, dir: ndir, cost: cst + add})
		}

		backx, backy := DX[(dir+2)%4], DY[(dir+2)%4]
		turns := []int{(dir + 1) % 4, (dir + 3) % 4}
		check(x+backx, y+backy, cost, 1, dir)
		for _, t := range turns {
			check(x, y, cost, TURN_COST, t)
		}
	}
	return dis
}

func main() {
	grid, sx, sy, tx, ty := parseInput()
	n, m = len(grid), len(grid[0])

	source := dijkstra(grid, sx, sy, []int{1})
	sink := rev_dijkstra(grid, tx, ty, []int{0, 1, 2, 3})
	best, cnt := INF, 0
	for d := 0; d < 4; d++ {
		if source[tx][ty][d] < best {
			best = source[tx][ty][d]
		}
	}

	for x := range source {
		for y := range source[x] {
			can := false
			for d := 0; d < 4; d++ {
				if source[x][y][d]+sink[x][y][d] == best {
					can = true
				}
			}
			if can {
				cnt++
			}
		}
	}
	fmt.Println(best, cnt)
}
