package main

import (
	"fmt"
	"log"
	"os"
	"slices"
)

func p1(a, b []int) int{
	ans := 0
	slices.Sort(a)
	slices.Sort(b)
	for i, _ := range a {
		d := a[i] - b[i]
		if d < 0 {d = -d}
		ans += d
	}
	return ans
}

func p2(a, b []int) int {
	f := make(map[int]int)
	ans := 0
	for _, v := range b {
		f[v] += 1
	}
	for _, v := range a {
		ans += v * f[v]
	}
	return ans
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {log.Fatal("Error opening input file.")}
	defer f.Close()

	var a, b []int
	for {
		var x, y int
		_, err := fmt.Fscan(f, &x, &y)
		if err != nil {break;}
		a = append(a, x)
		b = append(b, y)
	}

	fmt.Println(p1(a, b))
	fmt.Println(p2(a, b))
}