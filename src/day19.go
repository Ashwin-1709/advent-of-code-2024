package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	INPUT_FILE = "input.txt"
)

func parseInput() ([]string, []string) {
	f, _ := os.Open(INPUT_FILE)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var towels []string
	var patternList string
	scanner.Scan()
	patternList = scanner.Text()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			towels = append(towels, line)
		}
	}
	var patterns []string
	for _, pattern := range strings.Split(patternList, ",") {
		patterns = append(patterns, strings.TrimSpace(pattern))
	}
	return patterns, towels
}

func matchWays(towel string, patterns []string) int {
	dp := make(map[int]int)
	var f func(i int) int
	f = func(i int) int {
		if i == len(towel) {
			return 1
		}

		_, ok := dp[i]
		if ok {
			return dp[i]
		}

		dp[i] = 0
		left := len(towel) - i
		for _, pattern := range patterns {
			patternLen := len(pattern)
			if patternLen > left {
				continue
			}
			match := towel[i : i+patternLen]
			if match == pattern {
				dp[i] += f(i + patternLen)
			}
		}
		return dp[i]
	}

	return f(0)
}

func main() {
	patterns, towels := parseInput()
	cnt, ways := 0, 0
	for _, towel := range towels {
		match := matchWays(towel, patterns)
		ways += match
		if match > 0 {
			cnt++
		}
	}
	fmt.Println(cnt, ways)
}
