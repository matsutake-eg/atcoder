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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)

	ans := math.MaxInt64
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			s := i * j
			if s > n {
				break
			}
			ans = min(ans, n-s+abs(i-j))
		}
	}
	fmt.Println(ans)
}
