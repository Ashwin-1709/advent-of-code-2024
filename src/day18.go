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
	TURN_COST  = 1000
	ROWS       = 71
	COLS       = 71
	T          = 1024
)

var (
	n  = ROWS
	m  = COLS
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

type Pair struct {
	x int
	y int
}

func parseInput() []Pair {
	f, _ := os.Open(INPUT_FILE)
	defer f.Close()
	var pos []Pair
	reader := bufio.NewReader(f)
	for {
		var x, y int
		_, err := fmt.Fscanf(reader, "%d,%d\n", &y, &x)
		if err != nil {
			break
		}
		pos = append(pos, Pair{x, y})
	}
	return pos
}

func ok(x, y int) bool {
	return x >= 0 && x < ROWS && y >= 0 && y < COLS
}

func main() {
	pos := parseInput()

	bfs := func(sx, sy, tx, ty, time int) int {
		blocked := make([][]bool, n)
		dis := make([][]int, n)
		for x := range blocked {
			blocked[x] = make([]bool, m)
			dis[x] = make([]int, m)
			for y := range dis[x] {
				dis[x][y] = -1
			}
		}

		for i := 0; i < time; i++ {
			blocked[pos[i].x][pos[i].y] = true
		}

		Q := make(chan Pair, n*m)
		dis[sx][sy] = 0
		Q <- Pair{sx, sy}
		for len(Q) > 0 {
			cur := <-Q
			for i := range dx {
				nx, ny := cur.x+dx[i], cur.y+dy[i]
				if ok(nx, ny) && !blocked[nx][ny] && dis[nx][ny] == -1 {
					dis[nx][ny] = dis[cur.x][cur.y] + 1
					Q <- Pair{nx, ny}
				}
			}
		}

		return dis[tx][ty]
	}

	l, r, ans := 1, len(pos), -1
	for l <= r {
		mid := (l + r) / 2
		if bfs(0, 0, n-1, m-1, mid) == -1 {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	fmt.Println(pos[ans-1].y, pos[ans-1].x)
}
