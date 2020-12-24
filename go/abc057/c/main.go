package main

import (
	"fmt"
	"math"
	"strconv"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n int
	fmt.Scan(&n)

	ans := math.MaxInt64
	for i := 1; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		f := max(len(strconv.Itoa(i)), len(strconv.Itoa(n/i)))
		ans = min(ans, f)
	}
	fmt.Println(ans)
}
