package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	INPUT_FILE = "input.txt"
)

func get(cliques []string) string {
	slices.Sort(cliques)
	return strings.Join(cliques, ",")
}

func parseInput() map[string][]string {
	f, _ := os.Open(INPUT_FILE)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	graph := make(map[string][]string)

	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), "-")
		from, to := nodes[0], nodes[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	return graph
}

func getTrianglesWithNode(graph map[string][]string, node string) []string {
	var triangles []string
	for _, x := range graph[node] {
		for _, y := range graph[x] {
			if y == node {
				continue
			}
			if slices.Contains(graph[node], y) {
				triangles = append(triangles, get([]string{node, x, y}))
			}
		}
	}
	return triangles
}

func getTriangles(graph map[string][]string) int {
	cliques := make(map[string]bool)
	for k := range graph {
		if k[0] == 't' {
			for _, clique := range getTrianglesWithNode(graph, k) {
				cliques[clique] = true
			}
		}
	}

	return len(cliques)
}

func getCliqueOfSize(graph map[string][]string, node string, size int) string {
	if size == 1 {
		return node
	}
	var cliques []string
	var dfs func(node string, has map[string]bool)
	dfs = func(node string, has map[string]bool) {
		if len(cliques) == size {
			return
		}
		if len(has) == size-1 {
			has[node] = true
			for k := range has {
				cliques = append(cliques, k)
			}
			return
		}

		has[node] = true
		for _, x := range graph[node] {
			_, ok := has[x]
			if ok {
				continue
			}
			connected := true
			for k := range has {
				if !slices.Contains(graph[x], k) {
					connected = false
					break
				}
			}
			if connected {
				dfs(x, has)
			}
		}
	}

	dfs(node, make(map[string]bool))
	return get(cliques)
}

func main() {
	graph := parseInput()
	fmt.Println(getTriangles(graph))

	for size := 20; size >= 1; size-- {
		found := false
		for k := range graph {
			clique := getCliqueOfSize(graph, k, size)
			if len(clique) > 0 {
				fmt.Println(clique)
				found = true
				break
			}
		}
		if found {
			break
		}
	}
}
