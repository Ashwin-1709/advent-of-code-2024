package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

const (
	INPUT_FILE = "input.txt"
)

func parseInput() ([][]int, [][]int) {
	f, _ := os.Open(INPUT_FILE)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var keys, locks [][]int
	for {
		var grid []string
		read := false
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				break
			}
			grid = append(grid, line)
			read = true
		}
		if read == false {
			break
		}

		var holes, blocks []int
		for y := 0; y < len(grid[0]); y++ {
			cnt, block := 0, 0
			for x := 0; x < len(grid); x++ {
				if grid[x][y] == '.' {
					cnt++
				} else {
					block++
				}
			}
			blocks = append(blocks, block)
			holes = append(holes, cnt)
		}

		switch grid[0][0] {
		case '#':
			locks = append(locks, holes)
		default:
			keys = append(keys, blocks)
		}
	}
	return keys, locks
}

func get(lockId int, keys, locks [][]int, c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	match := 0
	for _, key := range keys {
		good := true
		for i := range key {
			good = (good && (locks[lockId][i] >= key[i]))
		}
		if good {
			match++
		}
	}
	c <- match
}

func main() {
	keys, locks := parseInput()

	c := make(chan int, len(locks))
	var wg sync.WaitGroup

	for lockId := 0; lockId < len(locks); lockId++ {
		wg.Add(1)
		go get(lockId, keys, locks, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	total := 0
	for match := range c {
		total += match
	}
	fmt.Println(total)
}
