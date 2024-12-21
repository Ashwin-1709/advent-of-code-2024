package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	NO_SOLUTION     = -1
	INPUT_FILE      = "input.txt"
	SIMULATION_TIME = 100
	COLS            = 101
	ROWS            = 103
)

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

type Pair struct {
	x int
	y int
}

func parseInput() []Robot {
	f, _ := os.Open(INPUT_FILE)
	defer f.Close()
	reader := bufio.NewReader(f)
	var robots []Robot
	for {
		var x, y, vx, vy int
		_, err := fmt.Fscanf(reader, "p=%d,%d v=%d,%d\n", &x, &y, &vx, &vy)
		if err != nil {
			break
		}
		robots = append(
			robots,
			Robot{
				x:  y,
				y:  x,
				vx: vy,
				vy: vx,
			},
		)
	}
	return robots
}

func get(val int, bound int) int {
	if val >= 0 {
		return val % bound
	}
	mod := (-val) % bound
	return (bound - mod) % bound
}

func simulate(robots []Robot, t int) []Robot {
	var upd []Robot
	for _, robot := range robots {
		curx, cury := robot.x+robot.vx*t, robot.y+robot.vy*t
		robot.x, robot.y = get(curx, ROWS), get(cury, COLS)
		upd = append(upd, robot)
	}
	return upd
}

func cost(freq map[Pair]int) int {
	ans := 1
	quad := make([]int, 4)
	for k, v := range freq {
		if k.x == ROWS/2 || k.y == COLS/2 {
			continue
		}
		if k.x < ROWS/2 {
			if k.y < COLS/2 {
				quad[0] += v
			} else {
				quad[2] += v
			}
		} else {
			if k.y < COLS/2 {
				quad[1] += v
			} else {
				quad[3] += v
			}
		}
	}

	for i := range quad {
		ans *= quad[i]
	}
	return ans
}

func print(robots []Robot) {
	block := make(map[Pair]bool)
	for _, robot := range robots {
		block[Pair{x: robot.x, y: robot.y}] = true
	}
	for x := 0; x < ROWS; x++ {
		str := ""
		for y := 0; y < COLS; y++ {
			_, ok := block[Pair{x: x, y: y}]
			if ok {
				str += "#"
			} else {
				str += "."
			}
		}
		fmt.Println(str)
	}
}

func solve(robots []Robot) {
	upd := simulate(robots, SIMULATION_TIME)
	freq := make(map[Pair]int)
	for _, robot := range upd {
		freq[Pair{x: robot.x, y: robot.y}]++
	}
	fmt.Println(cost(freq))
}

func main() {
	robots := parseInput()
	solve(robots)
	for T := 1; T <= 10000; T++ {
		upd := simulate(robots, T)
		fmt.Println(T)
		print(upd)
		fmt.Fprintf(os.Stderr, "Done: %d\n", T)
	}
}
