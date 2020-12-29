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
	var x, y int
	fmt.Scan(&x, &y)

	ans := math.MaxInt64
	if x <= y {
		ans = min(ans, y-x)
	}
	if a := -x; a <= y {
		ans = min(ans, y-a+1)
	}
	if b := -y; x <= b {
		ans = min(ans, b-x+1)
	}
	if a, b := -x, -y; a <= b {
		ans = min(ans, b-a+2)
	}
	fmt.Println(ans)
}
