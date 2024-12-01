package main

import (
	"fmt"
	"log"
	"os"
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
	f := make(map[int]int)
	for _, val := range b {
		f[val] += 1
	}

	for _, val := range a {
		ans += val * f[val]
	}
	fmt.Println(ans)
}
