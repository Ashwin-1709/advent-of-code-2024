package main

import (
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}

	defer file.Close()

	var a, b []int
	for {
		var x, y int
		_, err := fmt.Fscan(file, &x, &y)
		if err != nil {
			break
		}
		a = append(a, x)
		b = append(b, y)
	}

	ans := 0
	slices.Sort(a)
	slices.Sort(b)
	for i, _ := range a {
		diff := b[i] - a[i]
		if diff < 0 {
			diff = -diff
		}
		ans += diff
	}
	fmt.Println(ans)
}
