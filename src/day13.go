package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type delta struct {
	x int
	y int
}

type claw struct {
	buttonA delta
	buttonB delta
	prizeX  int
	prizeY  int
}

const (
	NO_SOLUTION    = -1
	INPUT_FILE     = "input.txt"
	SCALING_FACTOR = 10000000000000
)

func parseInput() []claw {
	f, _ := os.Open(INPUT_FILE)
	defer f.Close()
	reader := bufio.NewReader(f)
	var claws []claw

	line := false
	for {
		var ax, ay, bx, by, px, py int
		_, err := fmt.Fscanf(reader, "Button A: X+%d, Y+%d\n", &ax, &ay)
		if err != nil {
			if !line {
				line = true
				fmt.Fscanf(reader, "\n")
				continue
			} else {
				break
			}
		}
		fmt.Fscanf(reader, "Button B: X+%d, Y+%d\n", &bx, &by)
		fmt.Fscanf(reader, "Prize: X=%d, Y=%d\n", &px, &py)
		line = false
		claws = append(
			claws,
			claw{
				buttonA: delta{ax, ay},
				buttonB: delta{bx, by},
				prizeX:  px,
				prizeY:  py,
			},
		)
	}
	return claws
}

func getBruteCost(c claw) int {
	var tokens []int
	for pushA := 0; pushA <= 100; pushA++ {
		for pushB := 0; pushB <= 100; pushB++ {
			if c.buttonA.x*pushA+c.buttonB.x*pushB == c.prizeX &&
				c.buttonA.y*pushA+c.buttonB.y*pushB == c.prizeY {
				tokens = append(tokens, pushB+3*pushA)
			}
		}
	}

	slices.Sort(tokens)
	if len(tokens) == 0 {
		return NO_SOLUTION
	}

	return tokens[0]
}

func getScaledCost(c claw) int {
	/*
	   ax * u + bx * v == px
	   ay * ax * u + ay * bx * v == ay * px
	   ax * ay * u + ax * by * v == ax * py
	   v = (ay * px - ax * py) / (ay * bx - ax * by)
	   u = (px - bx * v) / ax
	*/
	c.prizeX += SCALING_FACTOR
	c.prizeY += SCALING_FACTOR
	num := (c.buttonA.y*c.prizeX - c.buttonA.x*c.prizeY)
	den := (c.buttonA.y*c.buttonB.x - c.buttonA.x*c.buttonB.y)
	if den == 0 || num%den != 0 {
		return NO_SOLUTION
	}
	pushB := num / den
	numA := (c.prizeX - c.buttonB.x*pushB)
	if numA%c.buttonA.x != 0 {
		return NO_SOLUTION
	}
	return (pushB + (numA*3)/c.buttonA.x)
}

func main() {
	claws := parseInput()
	part_1, part_2 := 0, 0
	for _, claw := range claws {
		cost, scaledCost := getBruteCost(claw), getScaledCost(claw)
		if cost != NO_SOLUTION {
			part_1 += cost
		}
		if scaledCost != NO_SOLUTION {
			part_2 += scaledCost
		}
	}

	fmt.Println(part_1, part_2)
}
