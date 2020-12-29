package main

import (
	"fmt"
	"math"
	"strconv"
)

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
	for i := 1; i <= n/2; i++ {
		sum := 0
		for _, r := range strconv.Itoa(i) {
			sum += int(r - '0')
		}
		for _, r := range strconv.Itoa(n - i) {
			sum += int(r - '0')
		}
		ans = min(ans, sum)
	}
	fmt.Println(ans)
}
