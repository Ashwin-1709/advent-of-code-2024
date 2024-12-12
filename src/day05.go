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

func atoi(s string) int {
	val, err := strconv.Atoi(s)
	handleError(err)
	return val
}

type Graph map[int][]int
type Adj map[int]map[int]bool

func parseInput() (Graph, [][]int, Adj) {
	f, err := os.Open("input.txt")
	handleError(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	readGraph := true
	var pages [][]int
	graph := make(Graph)
	adj := make(Adj)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			readGraph = false
			continue
		}

		switch readGraph {
		case true:
			nums := strings.Split(line, "|")
			from, to := atoi(nums[0]), atoi(nums[1])
			graph[from] = append(graph[from], to)
			if len(adj[from]) == 0 {
				adj[from] = make(map[int]bool)
			}
			adj[from][to] = true
		default:
			nums := strings.Split(line, ",")
			var page []int
			for _, num := range nums {
				page = append(page, atoi(num))
			}
			pages = append(pages, page)
		}
	}

	return graph, pages, adj
}

func good(page []int, a Adj) bool {
	for x := 0; x < len(page)-1; x++ {
		for y := x + 1; y < len(page); y++ {
			if !ok(page, a, x, y) {
				return false
			}
		}
	}
	return true
}

func ok(page []int, a Adj, x, y int) bool {
	l, r := page[x], page[y]
	_, ok := a[r][l]
	return !ok
}

func reachable(g Graph, src, dest int) bool {
	seen := make(map[int]bool)
	valid := true
	var dfs func(node int)
	dfs = func(node int) {
		fmt.Println(node)
		if node == src {
			valid = false
			return
		}
		_, ok := seen[node]
		if ok {
			return
		}
		seen[node] = true
		for _, child := range g[node] {
			dfs(child)
		}
	}
	dfs(dest)
	return valid
}

func correct(page []int, a Adj) int {
	for x := 0; x < len(page)-1; x++ {
		for y := 0; y < len(page)-x-1; y++ {
			if !ok(page, a, y, y+1) {
				page[y], page[y+1] = page[y+1], page[y]
			}
		}
	}
	return page[len(page)/2]
}

func main() {
	graph, pages, adj := parseInput()
	part_1, part_2 := 0, 0
	for _, page := range pages {
		if good(page, adj) {
			part_1 += page[(len(page)-1)/2]
		} else {
			part_2 += correct(page, adj)
		}
	}
	fmt.Println(part_1, part_2)
}
