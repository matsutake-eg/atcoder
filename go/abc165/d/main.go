package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var a, b, n float64
	fmt.Scan(&a, &b, &n)

	x := float64(min(int(b-1), int(n)))
	fmt.Println(int(math.Floor(a*x/b)) - int(a)*int(math.Floor(x/b)))
}
