package main

import (
	"fmt"
	"math"
	"sort"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	xs := make([]int, 3)
	r1 := 0.0
	for i := range xs {
		fmt.Scan(&xs[i])
		r1 += float64(xs[i])
	}
	sort.Ints(xs)
	r2 := 0.0
	if xs[0]+xs[1] < xs[2] {
		r2 = float64(xs[2] - xs[1] - xs[0])
	}
	fmt.Println((r1*r1 - r2*r2) * math.Pi)
}
