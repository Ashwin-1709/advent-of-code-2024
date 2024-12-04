package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type pair struct {
	l int
	r int
}

func parseInput() string {
	f, err := os.Open("input.txt")
	handleError(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	command := ""
	for scanner.Scan() {
		line := scanner.Text()
		command += line
	}
	return command
}

func extractMuls(command string) ([]pair, []pair) {
	var muls, enabledMuls []pair
	pattern := regexp.MustCompile(`(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`)
	integerPattern := regexp.MustCompile(`\d+`)
	matches := pattern.FindAllString(command, -1)
	enabled := true
	
	for _, match := range matches {
		switch match {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			vals := integerPattern.FindAllStringSubmatch(match, -1)
			var p pair
			var err error

			p.l, err = strconv.Atoi(vals[0][0])
			handleError(err)

			p.r, err = strconv.Atoi(vals[1][0])
			handleError(err)

			muls = append(muls, p)
			if enabled {
				enabledMuls = append(enabledMuls, p)
			}
		}
	}
	return muls, enabledMuls
}

func compute(muls []pair) int {
	ans := 0
	for _, p := range muls {
		ans += p.l * p.r
	}
	return ans
}

func main() {
	command := parseInput()
	muls, enabledMuls := extractMuls(command)
	fmt.Println(compute(muls), compute(enabledMuls))
}
