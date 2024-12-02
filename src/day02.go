package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInput(path string) [][]int {
	file, err := os.Open(path)
	checkError(err)
	scanner := bufio.NewScanner(file)
	var reports [][]int
	for scanner.Scan() {
		report := scanner.Text()
		var row []int
		for _, level := range strings.Fields(report) {
			val, err := strconv.Atoi(level)
			checkError(err)
			row = append(row, val)
		}
		reports = append(reports, row)
	}
	return reports
}

func validAdjacent(x, y int) bool {
	distance := x - y
	return ((distance <= -1 && distance >= -3) || (distance >= 1 && distance <= 3))
}

func getPredicate(x, y int) func(x, y int) bool {
	if x > y {
		return func(x, y int) bool { return (x > y && validAdjacent(x, y)) }
	} else {
		return func(x, y int) bool { return (y > x && validAdjacent(x, y)) }
	}
}

func isSafe(report []int) bool {
	if len(report) <= 1 {
		return true
	}
	check := getPredicate(report[0], report[1])
	for i := 1; i < len(report); i++ {
		if !check(report[i-1], report[i]) {
			return false
		}
	}
	return true
}

func safeCount(reports [][]int) int {
	safe := 0
	for _, report := range reports {
		if isSafe(report) {
			safe++
		}
	}
	return safe
}

func almostSafeCount(reports [][]int) int {
	almost := 0
	for _, report := range reports {
		canSave := isSafe(report)
		for id, _ := range report {
			subset := slices.Clone(report)
			subset = slices.Delete(subset, id, id+1)
			if canSave || isSafe(subset) {
				canSave = true
				break
			}
		}
		if canSave {
			almost++
		}
	}
	return almost
}

func main() {
	reports := parseInput("input.txt")
	fmt.Println(safeCount(reports), almostSafeCount(reports))
}
