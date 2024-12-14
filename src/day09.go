package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type info struct {
	fileCount  int
	spaceCount int
	id         int
	moved      bool
}

func parseInput() string {
	f, err := os.Open("input.txt")
	handleError(err)
	defer f.Close()
	scanner := bufio.NewReader(f)
	var s string
	_, e := fmt.Fscan(scanner, &s)
	handleError(e)
	return s
}

func individualFit(a []info) []int {
	var sorted []int
	mod := make([]info, len(a))
	copy(mod, a)

	l, r := 0, len(a)-1
	for {
		if r < l {
			break
		}
		for f := 0; f < mod[l].fileCount; f++ {
			sorted = append(sorted, mod[l].id)
		}
		spaces := mod[l].spaceCount
		for {
			if spaces == 0 || r <= l {
				break
			}
			can := mod[r].fileCount
			if can < spaces {
				spaces -= can
				for f := 0; f < can; f++ {
					sorted = append(sorted, mod[r].id)
				}
				mod[r].fileCount = 0
				r--
			} else {
				for f := 0; f < spaces; f++ {
					sorted = append(sorted, mod[r].id)
				}
				mod[r].fileCount -= spaces
				spaces = 0
			}
		}
		l++
	}

	return sorted
}

func groupFit(a []info) []int {
	var sorted []int
	mod := make([]info, len(a))
	copy(mod, a)

	i := len(mod) - 1
	for {
		if i == 0 {
			break
		}
		found := false
		for j := 0; j < i; j++ {
			if mod[j].spaceCount >= mod[i].fileCount {
				diff := mod[j].spaceCount - mod[i].fileCount
				last := mod[i]
				copy(mod[j+1:], mod[j:i])
				mod[j+1] = last
				mod[j].spaceCount = 0
				mod[j+1].spaceCount = diff
				mod[i].spaceCount += last.fileCount + last.spaceCount
				mod[i].moved = true
				found = true
				break
			}
		}
		if !found {
			i--
		}
	}

	for i := range mod {
		for j := 0; j < mod[i].fileCount; j++ {
			sorted = append(sorted, mod[i].id)
		}
		for j := 0; j < mod[i].spaceCount; j++ {
			sorted = append(sorted, 0)
		}
	}

	return sorted
}

func compute(sorted []int) int {
	ans := 0
	for i := range sorted {
		ans += i * sorted[i]
	}
	return ans
}

func main() {
	s := parseInput()
	var a []info
	for i := 0; i < len(s); i += 2 {
		fileCount, _ := strconv.Atoi(string(s[i]))
		spaceCount := 0
		if i+1 < len(s) {
			nxt, _ := strconv.Atoi(string(s[i+1]))
			spaceCount += nxt
		}
		a = append(a, info{fileCount, spaceCount, i / 2, false})
	}

	part_1, part_2 := individualFit(a), groupFit(a)
	fmt.Println(compute(part_1), compute(part_2))
}
