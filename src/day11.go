package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput() []int {
	f, _ := os.Open("input.txt")
	defer f.Close()
	reader := bufio.NewReader(f)
	var stones []int
	for {
		var stone int
		_, ok := fmt.Fscan(reader, &stone)
		if ok != nil {
			break
		}
		stones = append(stones, stone)
	}
	return stones
}

type state struct {
	stone int
	turns int
}

var dp = make(map[state]int)

func f(stone, turns int) int {
	cur := state{stone, turns}
	val, cached := dp[cur]
	if cached {
		return val
	}
	if turns == 0 {
		dp[cur] = 1
		return 1
	}

	if stone == 0 {
		dp[cur] = f(1, turns-1)
	} else {
		str := strconv.Itoa(stone)
		if len(str)%2 != 0 {
			dp[cur] = f(2024*stone, turns-1)
		} else {
			l, r := str[:len(str)/2], str[len(str)/2:]
			lv, _ := strconv.Atoi(l)
			rv, _ := strconv.Atoi(r)
			dp[cur] = f(lv, turns-1) + f(rv, turns-1)
		}
	}
	return dp[cur]
}

func simulate(stones []int, k int) []int {
	for i := 0; i < k; i++ {
		var nxt []int
		for _, stone := range stones {
			if stone == 0 {
				nxt = append(nxt, 1)
			} else {
				str := strconv.Itoa(stone)
				switch len(str) % 2 {
				case 0:
					l, r := str[:len(str)/2], str[len(str)/2:]
					lv, _ := strconv.Atoi(l)
					rv, _ := strconv.Atoi(r)
					nxt = append(nxt, lv, rv)
				default:
					nxt = append(nxt, 2024*stone)
				}
			}
		}
		stones = nxt
	}
	return stones
}

func main() {
	stones := parseInput()
	ans := 0
	for _, stone := range stones {
		ans += f(stone, 75)
	}

	fmt.Println(ans)
}
