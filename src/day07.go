package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInput() ([]int, [][]int) {
	f, err := os.Open("input.txt")
	handleError(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lhs []int
	var rhs [][]int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		left, err := strconv.Atoi(parts[0])
		handleError(err)
		lhs = append(lhs, left)

		right := strings.TrimSpace(parts[1])
		list := strings.Split(right, " ")
		var ops []int
		for _, str := range list {
			num, err := strconv.Atoi(str)
			handleError(err)
			ops = append(ops, num)
		}
		rhs = append(rhs, ops)
	}
	return lhs, rhs
}

func concat(x, y int) int {
	l := strconv.Itoa(x)
	r := strconv.Itoa(y)
	z := l + r
	val, err := strconv.Atoi(z)
	handleError(err)
	return val
}

func check(lhs int, rhs []int) bool {
	for mask := 0; mask < 1<<(len(rhs)-1); mask++ {
		val := rhs[0]
		for bit := 0; bit < len(rhs)-1; bit++ {
			if (mask >> bit & 1) != 0 {
				val = val + rhs[bit+1]
			} else {
				val = val * rhs[bit+1]
			}
		}
		if val == lhs {
			return true
		}
	}
	return false
}

func ternary(lhs int, rhs []int) bool {
	can := false
	var f func(cur, id int)
	f = func(cur, id int) {
		if can {
			return
		}
		if id == len(rhs)-1 {
			if cur == lhs {
				can = true
			}
			return
		}

		nxt := rhs[id+1]
		f(cur+nxt, id+1)
		f(cur*nxt, id+1)
		f(concat(cur, nxt), id+1)
	}
	f(rhs[0], 0)
	return can
}

func main() {
	lhs, rhs := parseInput()
	part_1, part_2 := 0, 0
	for i, _ := range lhs {
		if check(lhs[i], rhs[i]) {
			part_1 += lhs[i]
		}
		if ternary(lhs[i], rhs[i]) {
			part_2 += lhs[i]
		}
	}
	fmt.Println(part_1, part_2)
}
